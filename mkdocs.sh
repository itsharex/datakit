#!/bin/bash
# author: tanbiao
# date: Fri Jun 24 10:59:21 CST 2022
#
# This tool used to generate & publish datakit related docs to docs.guance.com.
#

RED="\033[31m"
GREEN="\033[32m"
YELLOW="\033[33m"
CLR="\033[0m"

mkdocs_dir=~/git/dataflux-doc
lang=zh
port=8000
bind=0.0.0.0

usage() {
	echo "" 1>&2;
	echo "mkdocs.sh used to build/preview/release DataKit documents." 1>&2;
	echo "" 1>&2;
	echo "Usage: " 1>&2;
	echo "" 1>&2;
	echo "  ./mkdocs.sh -V string: Set version, such as 1.2.3" 1>&2;
	echo "              -D string: Set workdir, such as my-test" 1>&2;
	echo "              -B: Do not build datakit" 1>&2;
	echo "              -C: Check exported docs" 1>&2;
	echo "              -L: Specify language(zh/en)" 1>&2;
	echo "              -C: check(lint) generated docs" 1>&2;
	echo "              -p: Specify local port(default 8000)" 1>&2;
	echo "              -b: Specify local bind(default 0.0.0.0)" 1>&2;
	echo "              -h: Show help" 1>&2;
	echo "" 1>&2;
	exit 1;
}

while getopts "V:D:L:p:b:BCh" arg; do
	case "${arg}" in
		V)
			version="${OPTARG}"
			;;
		L)
		 lang="${OPTARG}"
		 ;;

		D)
			mkdocs_dir="${OPTARG}"
			printf "${YELLOW}> Set workdir to '%s'${CLR}\n" $mkdocs_dir
			;;

		C)
			check_doc=true;
			;;

		B)
			no_build=true;
			;;

		C)
			do_check=true;
			;;

		h)
			usage
			;;

		p)
			port="${OPTARG}"
			;;

		b)
			bind="${OPTARG}"
			;;

		*)
			echo "invalid args $@";
			usage
			;;
	esac
done
shift $((OPTIND-1))

# detect workdir
if [ ! -d $mkdocs_dir ]; then
	mkdir -p ${mkdocs_dir}/docs/{datakit,developers}
fi

# if -v not set...
if [ -z $version ]; then
	# get online datakit version
	latest_version=$(curl https://static.guance.com/datakit/version | grep '"version"' | awk -F'"' '{print $4}')

	printf "${YELLOW}> Version missing, use latest version '%s'${CLR}\n" $latest_version
	version="${latest_version}"
fi

tmp_doc_dir=.docs
base_docs_dir=${mkdocs_dir}/docs

######################################
# list i18n languages
######################################
i18n=(
	"zh"
	"en"
	# add more...
)

######################################
# prepare workdirs
######################################
# clear tmp dir
rm -rf $tmp_doc_dir/*
# create workdirs
for _lang in "${i18n[@]}"; do
	mkdir -p $base_docs_dir/${_lang}/datakit \
		$base_docs_dir/${_lang}/developers \
		$tmp_doc_dir/${_lang}
	done

######################################
# select datakit binary
######################################
arch=$(uname -m)
if [[ "$arch" == "x86_64" ]]; then
	arch=amd64
else
	arch=arm64
fi

if [[ "$OSTYPE" == "darwin"* ]]; then
	os="darwin"
	datakit=dist/datakit-${os}-${arch}/datakit
elif [[ "$OSTYPE" == "linux-gnu"* ]]; then
	os="linux"
	datakit=dist/datakit-${os}-${arch}/datakit
else              # if under windows(amd64):
	datakit=datakit # windows 下应该设置了对应的 $PATH
fi

if [[ ! $no_build ]]; then
	printf "${GREEN}> Building datakit...${CLR}\n"
	make || exit -1
fi

######################################
# export all docs to temp dir
######################################
printf "${GREEN}> Export internal docs to %s${CLR}\n" $tmp_doc_dir
truncate -s 0 .mkdocs.log
LOGGER_PATH=.mkdocs.log $datakit doc \
	--export-docs $tmp_doc_dir \
	--ignore demo \
	--version "${version}"

if [ $? -ne 0 ]; then
	printf "${RED}[E] Export docs failed${CLR}\n"
	exit -1
fi

######################################
# copy docs to different mkdocs sub-dirs
######################################
printf "${GREEN}> Copy docs...${CLR}\n"
for _lang in "${i18n[@]}"; do
	# copy .pages
	printf "${GREEN}> Copy pages(%s) to repo datakit ...${CLR}\n" $_lang
	cp man/docs/$_lang/datakit.pages $base_docs_dir/$_lang/datakit/.pages

	# move specific docs to developers
	printf "${GREEN}> Copy docs(%s) to repo developers ...${CLR}\n" $_lang
	cp $tmp_doc_dir/${_lang}/pythond.md                ${base_docs_dir}/$_lang/developers
	cp $tmp_doc_dir/${_lang}/pipeline.md               ${base_docs_dir}/$_lang/developers
	cp $tmp_doc_dir/${_lang}/datakit-pl-global.md      ${base_docs_dir}/$_lang/developers
	cp $tmp_doc_dir/${_lang}/datakit-pl-how-to.md      ${base_docs_dir}/$_lang/developers
	cp $tmp_doc_dir/${_lang}/datakit-refer-table.md    ${base_docs_dir}/$_lang/developers

	# copy specific docs to datakit
	printf "${GREEN}> Copy docs(%s) to repo datakit ...${CLR}\n" $_lang
	cp $tmp_doc_dir/${_lang}/*.md $base_docs_dir/${_lang}/datakit/

	# NOTE: Only check Chinese documents.
	if [[ $check_doc && ${_lang} == "zh" ]]; then
		printf "${GREEN}> markdownlint %s...${CLR}\n" $base_docs_dir/${_lang}/datakit
		markdownlint $base_docs_dir/${_lang}/datakit 2>&1 | tee md.lint
		if [ -s md.lint ]; then
			exit -1;
		fi

		printf "${GREEN}> cspell %s...${CLR}\n" $base_docs_dir/${_lang}/datakit
		cspell lint -c cspell/cspell.json --no-progress $base_docs_dir/${_lang}/datakit | tee cspell.lint
		if [ -s cspell.lint ]; then
			exit -1;
		fi

		printf "${GREEN}> mdm %s...${CLR}\n" $base_docs_dir/${_lang}/datakit
		CGO_CFLAGS="-Wno-undef-prefix -Wno-deprecated-declarations"	go run cmd/make/make.go -mdm $base_docs_dir/${_lang}/datakit &>/dev/null && \
			{ echo "\n------\n[E] Some bad docs got invalid format on Unicode/ASCII. See https://docs.guance.com/datakit/mkdocs-howto/#zh-en-mix\n"; exit -1; } || \
				{ echo 'Unicode/ASCII format ok.'; };
		exit 0;
	fi
done

######################################
# start mkdocs local server
######################################
printf "${GREEN}> Start mkdocs on ${bind}:${port}...${CLR}\n"
cd $mkdocs_dir &&
	mkdocs serve -f mkdocs.${lang}.yml -a ${bind}:${port}  2>&1 | tee mkdocs.log
