package grok

import (
	"fmt"
	"regexp"
)

type GrokRegexp struct {
	Pattern             string
	DenormalizedPattern string
	Re                  *regexp.Regexp
}

func (g *GrokRegexp) Run(content interface{}) (map[string]string, error) {
	if g.Re == nil {
		return nil, fmt.Errorf("not complied")
	}
	result := map[string]string{}

	switch v := content.(type) {
	case []byte:
		match := g.Re.FindSubmatch(v)
		if len(match) == 0 {
			return nil, fmt.Errorf("no match")
		}
		for i, name := range g.Re.SubexpNames() {
			if name != "" {
				result[name] = string(match[i])
			}
		}
	case string:
		match := g.Re.FindStringSubmatch(v)
		if len(match) == 0 {
			return nil, fmt.Errorf("no match")
		}
		for i, name := range g.Re.SubexpNames() {
			if name != "" {
				result[name] = match[i]
			}
		}
	}
	return result, nil
}

func CopyDefalutPatterns() map[string]string {
	ret := map[string]string{}
	for k, v := range defalutPatterns {
		ret[k] = v
	}
	return ret
}

var defalutPatterns = map[string]string{
	"USERNAME":             `[a-zA-Z0-9._-]+`,
	"USER":                 `%{USERNAME}`,
	"EMAILLOCALPART":       `[a-zA-Z][a-zA-Z0-9_.+-=:]+`,
	"EMAILADDRESS":         `%{EMAILLOCALPART}@%{HOSTNAME}`,
	"HTTPDUSER":            `%{EMAILADDRESS}|%{USER}`,
	"INT":                  `(?:[+-]?(?:[0-9]+))`,
	"BASE10NUM":            `([+-]?(?:[0-9]+(?:\.[0-9]+)?)|\.[0-9]+)`,
	"NUMBER":               `(?:%{BASE10NUM})`,
	"BASE16NUM":            `(0[xX]?[0-9a-fA-F]+)`,
	"POSINT":               `\b(?:[1-9][0-9]*)\b`,
	"NONNEGINT":            `\b(?:[0-9]+)\b`,
	"WORD":                 `\b\w+\b`,
	"NOTSPACE":             `\S+`,
	"SPACE":                `\s*`,
	"DATA":                 `.*?`,
	"GREEDYDATA":           `.*`,
	"QUOTEDSTRING":         `"([^"\\]*(\\.[^"\\]*)*)"|\'([^\'\\]*(\\.[^\'\\]*)*)\'`,
	"UUID":                 `[A-Fa-f0-9]{8}-(?:[A-Fa-f0-9]{4}-){3}[A-Fa-f0-9]{12}`,
	"MAC":                  `(?:%{CISCOMAC}|%{WINDOWSMAC}|%{COMMONMAC})`,
	"CISCOMAC":             `(?:(?:[A-Fa-f0-9]{4}\.){2}[A-Fa-f0-9]{4})`,
	"WINDOWSMAC":           `(?:(?:[A-Fa-f0-9]{2}-){5}[A-Fa-f0-9]{2})`,
	"COMMONMAC":            `(?:(?:[A-Fa-f0-9]{2}:){5}[A-Fa-f0-9]{2})`,
	"IPV6":                 `((([0-9A-Fa-f]{1,4}:){7}([0-9A-Fa-f]{1,4}|:))|(([0-9A-Fa-f]{1,4}:){6}(:[0-9A-Fa-f]{1,4}|((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3})|:))|(([0-9A-Fa-f]{1,4}:){5}(((:[0-9A-Fa-f]{1,4}){1,2})|:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3})|:))|(([0-9A-Fa-f]{1,4}:){4}(((:[0-9A-Fa-f]{1,4}){1,3})|((:[0-9A-Fa-f]{1,4})?:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(([0-9A-Fa-f]{1,4}:){3}(((:[0-9A-Fa-f]{1,4}){1,4})|((:[0-9A-Fa-f]{1,4}){0,2}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(([0-9A-Fa-f]{1,4}:){2}(((:[0-9A-Fa-f]{1,4}){1,5})|((:[0-9A-Fa-f]{1,4}){0,3}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(([0-9A-Fa-f]{1,4}:){1}(((:[0-9A-Fa-f]{1,4}){1,6})|((:[0-9A-Fa-f]{1,4}){0,4}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(:(((:[0-9A-Fa-f]{1,4}){1,7})|((:[0-9A-Fa-f]{1,4}){0,5}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:)))(%.+)?`,
	"IPV4":                 `(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)`,
	"IP":                   `(?:%{IPV6}|%{IPV4})`,
	"HOSTNAME":             `\b(?:[0-9A-Za-z][0-9A-Za-z-]{0,62})(?:\.(?:[0-9A-Za-z][0-9A-Za-z-]{0,62}))*(\.?|\b)`,
	"HOST":                 `%{HOSTNAME}`,
	"IPORHOST":             `(?:%{IP}|%{HOSTNAME})`,
	"HOSTPORT":             `%{IPORHOST}:%{POSINT}`,
	"PATH":                 `(?:%{UNIXPATH}|%{WINPATH})`,
	"UNIXPATH":             `(/[\w_%!$@:.,-]?/?)(\S+)?`,
	"TTY":                  `(?:/dev/(pts|tty([pq])?)(\w+)?/?(?:[0-9]+))`,
	"WINPATH":              `([A-Za-z]:|\\)(?:\\[^\\?*]*)+`,
	"URIPROTO":             `[A-Za-z]+(\+[A-Za-z+]+)?`,
	"URIHOST":              `%{IPORHOST}(?::%{POSINT:port})?`,
	"URIPATH":              `(?:/[A-Za-z0-9$.+!*'(){},~:;=@#%_\-]*)+`,
	"URIPARAM":             `\?[A-Za-z0-9$.+!*'|(){},~@#%&/=:;_?\-\[\]<>]*`,
	"URIPATHPARAM":         `%{URIPATH}(?:%{URIPARAM})?`,
	"URI":                  `%{URIPROTO}://(?:%{USER}(?::[^@]*)?@)?(?:%{URIHOST})?(?:%{URIPATHPARAM})?`,
	"MONTH":                `\b(?:Jan(?:uary|uar)?|Feb(?:ruary|ruar)?|M(?:a|ä)?r(?:ch|z)?|Apr(?:il)?|Ma(?:y|i)?|Jun(?:e|i)?|Jul(?:y)?|Aug(?:ust)?|Sep(?:tember)?|O(?:c|k)?t(?:ober)?|Nov(?:ember)?|De(?:c|z)(?:ember)?)\b`,
	"MONTHNUM":             `(?:0?[1-9]|1[0-2])`,
	"MONTHNUM2":            `(?:0[1-9]|1[0-2])`,
	"MONTHDAY":             `(?:(?:0[1-9])|(?:[12][0-9])|(?:3[01])|[1-9])`,
	"DAY":                  `(?:Mon(?:day)?|Tue(?:sday)?|Wed(?:nesday)?|Thu(?:rsday)?|Fri(?:day)?|Sat(?:urday)?|Sun(?:day)?)`,
	"YEAR":                 `(\d\d){1,2}`,
	"HOUR":                 `(?:2[0123]|[01]?[0-9])`,
	"MINUTE":               `(?:[0-5][0-9])`,
	"SECOND":               `(?:(?:[0-5]?[0-9]|60)(?:[:.,][0-9]+)?)`,
	"TIME":                 `([^0-9]?)%{HOUR}:%{MINUTE}(?::%{SECOND})([^0-9]?)`,
	"DATE_US":              `%{MONTHNUM}[/-]%{MONTHDAY}[/-]%{YEAR}`,
	"DATE_EU":              `%{MONTHDAY}[./-]%{MONTHNUM}[./-]%{YEAR}`,
	"ISO8601_TIMEZONE":     `(?:Z|[+-]%{HOUR}(?::?%{MINUTE}))`,
	"ISO8601_SECOND":       `(?:%{SECOND}|60)`,
	"TIMESTAMP_ISO8601":    `%{YEAR}-%{MONTHNUM}-%{MONTHDAY}[T ]%{HOUR}:?%{MINUTE}(?::?%{SECOND})?%{ISO8601_TIMEZONE}?`,
	"DATE":                 `%{DATE_US}|%{DATE_EU}`,
	"DATESTAMP":            `%{DATE}[- ]%{TIME}`,
	"TZ":                   `(?:[PMCE][SD]T|UTC)`,
	"DATESTAMP_RFC822":     `%{DAY} %{MONTH} %{MONTHDAY} %{YEAR} %{TIME} %{TZ}`,
	"DATESTAMP_RFC2822":    `%{DAY}, %{MONTHDAY} %{MONTH} %{YEAR} %{TIME} %{ISO8601_TIMEZONE}`,
	"DATESTAMP_OTHER":      `%{DAY} %{MONTH} %{MONTHDAY} %{TIME} %{TZ} %{YEAR}`,
	"DATESTAMP_EVENTLOG":   `%{YEAR}%{MONTHNUM2}%{MONTHDAY}%{HOUR}%{MINUTE}%{SECOND}`,
	"HTTPDERROR_DATE":      `%{DAY} %{MONTH} %{MONTHDAY} %{TIME} %{YEAR}`,
	"SYSLOGTIMESTAMP":      `%{MONTH} +%{MONTHDAY} %{TIME}`,
	"PROG":                 `[\x21-\x5a\x5c\x5e-\x7e]+`,
	"SYSLOGPROG":           `%{PROG:program}(?:\[%{POSINT:pid}\])?`,
	"SYSLOGHOST":           `%{IPORHOST}`,
	"SYSLOGFACILITY":       `<%{NONNEGINT:facility}.%{NONNEGINT:priority}>`,
	"HTTPDATE":             `%{MONTHDAY}/%{MONTH}/%{YEAR}:%{TIME} %{INT}`,
	"QS":                   `%{QUOTEDSTRING}`,
	"SYSLOGBASE":           `%{SYSLOGTIMESTAMP:timestamp} (?:%{SYSLOGFACILITY} )?%{SYSLOGHOST:logsource} %{SYSLOGPROG}:`,
	"COMMONAPACHELOG":      `%{IPORHOST:clientip} %{HTTPDUSER:ident} %{USER:auth} \[%{HTTPDATE:timestamp}\] "(?:%{WORD:verb} %{NOTSPACE:request}(?: HTTP/%{NUMBER:httpversion})?|%{DATA:rawrequest})" %{NUMBER:response} (?:%{NUMBER:bytes}|-)`,
	"COMBINEDAPACHELOG":    `%{COMMONAPACHELOG} %{QS:referrer} %{QS:agent}`,
	"HTTPD20_ERRORLOG":     `\[%{HTTPDERROR_DATE:timestamp}\] \[%{LOGLEVEL:loglevel}\] (?:\[client %{IPORHOST:clientip}\] ){0,1}%{GREEDYDATA:errormsg}`,
	"HTTPD24_ERRORLOG":     `\[%{HTTPDERROR_DATE:timestamp}\] \[%{WORD:module}:%{LOGLEVEL:loglevel}\] \[pid %{POSINT:pid}:tid %{NUMBER:tid}\]( \(%{POSINT:proxy_errorcode}\)%{DATA:proxy_errormessage}:)?( \[client %{IPORHOST:client}:%{POSINT:clientport}\])? %{DATA:errorcode}: %{GREEDYDATA:message}`,
	"HTTPD_ERRORLOG":       `%{HTTPD20_ERRORLOG}|%{HTTPD24_ERRORLOG}`,
	"LOGLEVEL":             `([Aa]lert|ALERT|[Tt]race|TRACE|[Dd]ebug|DEBUG|[Nn]otice|NOTICE|[Ii]nfo|INFO|[Ww]arn?(?:ing)?|WARN?(?:ING)?|[Ee]rr?(?:or)?|ERR?(?:OR)?|[Cc]rit?(?:ical)?|CRIT?(?:ICAL)?|[Ff]atal|FATAL|[Ss]evere|SEVERE|EMERG(?:ENCY)?|[Ee]merg(?:ency)?)`,
	"COMMONENVOYACCESSLOG": `\[%{TIMESTAMP_ISO8601:timestamp}\] \"%{DATA:method} (?:%{URIPATH:uri_path}(?:%{URIPARAM:uri_param})?|%{DATA:}) %{DATA:protocol}\" %{NUMBER:status_code} %{DATA:response_flags} %{NUMBER:bytes_received} %{NUMBER:bytes_sent} %{NUMBER:duration} (?:%{NUMBER:upstream_service_time}|%{DATA:tcp_service_time}) \"%{DATA:forwarded_for}\" \"%{DATA:user_agent}\" \"%{DATA:request_id}\" \"%{DATA:authority}\" \"%{DATA:upstream_service}\"`,
}
