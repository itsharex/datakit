// Code generated by "stringer -type=TNSType"; DO NOT EDIT.

package go_ora

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[NCHAR-1]
	_ = x[NUMBER-2]
	_ = x[BInteger-3]
	_ = x[FLOAT-4]
	_ = x[NullStr-5]
	_ = x[VarNum-6]
	_ = x[LONG-8]
	_ = x[VARCHAR-9]
	_ = x[ROWID-11]
	_ = x[DATE-12]
	_ = x[VarRaw-15]
	_ = x[BFloat-21]
	_ = x[BDouble-22]
	_ = x[RAW-23]
	_ = x[LongRaw-24]
	_ = x[UINT-68]
	_ = x[LongVarChar-94]
	_ = x[LongVarRaw-95]
	_ = x[CHAR-96]
	_ = x[CHARZ-97]
	_ = x[IBFloat-100]
	_ = x[IBDouble-101]
	_ = x[REFCURSOR-102]
	_ = x[OCIXMLType-108]
	_ = x[XMLType-109]
	_ = x[OCIRef-110]
	_ = x[OCIClobLocator-112]
	_ = x[OCIBlobLocator-113]
	_ = x[OCIFileLocator-114]
	_ = x[ResultSet-116]
	_ = x[OCIString-155]
	_ = x[OCIDate-156]
	_ = x[TimeStampDTY-180]
	_ = x[TimeStampTZ_DTY-181]
	_ = x[IntervalYM_DTY-182]
	_ = x[IntervalDS_DTY-183]
	_ = x[TimeTZ-186]
	_ = x[TIMESTAMP-187]
	_ = x[TIMESTAMPTZ-188]
	_ = x[IntervalYM-189]
	_ = x[IntervalDS-190]
	_ = x[UROWID-208]
	_ = x[TimeStampLTZ_DTY-231]
	_ = x[TimeStampeLTZ-232]
}

const _OracleType_name = "NCHARNUMBERSB1FLOATNullStrVarNumLONGVARCHARROWIDDATEVarRawBFloatBDoubleRAWLongRawUINTLongVarCharLongVarRawCHARCHARZIBFloatIBDoubleREFCURSOROCIXMLTypeXMLTypeOCIRefOCIClobLocatorOCIBlobLocatorOCIFileLocatorResultSetOCIStringOCIDateTimeStampDTYTimeStampTZ_DTYIntervalYM_DTYIntervalDS_DTYTimeTZTIMESTAMPTimeStampTZIntervalYMIntervalDSUROWIDTimeStampLTZ_DTYTimeStampeLTZ"

var _OracleType_map = map[TNSType]string{
	1:   _OracleType_name[0:5],
	2:   _OracleType_name[5:11],
	3:   _OracleType_name[11:14],
	4:   _OracleType_name[14:19],
	5:   _OracleType_name[19:26],
	6:   _OracleType_name[26:32],
	8:   _OracleType_name[32:36],
	9:   _OracleType_name[36:43],
	11:  _OracleType_name[43:48],
	12:  _OracleType_name[48:52],
	15:  _OracleType_name[52:58],
	21:  _OracleType_name[58:64],
	22:  _OracleType_name[64:71],
	23:  _OracleType_name[71:74],
	24:  _OracleType_name[74:81],
	68:  _OracleType_name[81:85],
	94:  _OracleType_name[85:96],
	95:  _OracleType_name[96:106],
	96:  _OracleType_name[106:110],
	97:  _OracleType_name[110:115],
	100: _OracleType_name[115:122],
	101: _OracleType_name[122:130],
	102: _OracleType_name[130:139],
	108: _OracleType_name[139:149],
	109: _OracleType_name[149:156],
	110: _OracleType_name[156:162],
	112: _OracleType_name[162:176],
	113: _OracleType_name[176:190],
	114: _OracleType_name[190:204],
	116: _OracleType_name[204:213],
	155: _OracleType_name[213:222],
	156: _OracleType_name[222:229],
	180: _OracleType_name[229:241],
	181: _OracleType_name[241:256],
	182: _OracleType_name[256:270],
	183: _OracleType_name[270:284],
	186: _OracleType_name[284:290],
	187: _OracleType_name[290:299],
	188: _OracleType_name[299:310],
	189: _OracleType_name[310:320],
	190: _OracleType_name[320:330],
	208: _OracleType_name[330:336],
	231: _OracleType_name[336:352],
	232: _OracleType_name[352:365],
}

func (i TNSType) String() string {
	if str, ok := _OracleType_map[i]; ok {
		return str
	}
	return "TNSType(" + strconv.FormatInt(int64(i), 10) + ")"
}
