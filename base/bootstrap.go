package base

import (
	"io"
	"log"
	"log/syslog"
	"net/url"
)

type Bootstrap struct {
	name    string
	version Version
	logging *LoggingAdapter
}

func NewBootstrap(name string) *Bootstrap {
	return &Bootstrap{name: name}
}

func (b *Bootstrap) Version(major, minor, build int) *Bootstrap {
	b.version.Major = major
	b.version.Minor = minor
	b.version.Build = build

	return b
}

func (b *Bootstrap) Major(major int) *Bootstrap {
	b.version.Major = major

	return b
}

func (b *Bootstrap) Minor(minor int) *Bootstrap {
	b.version.Minor = minor

	return b
}

func (b *Bootstrap) Build(build int) *Bootstrap {
	b.version.Build = build

	return b
}

func (b *Bootstrap) Revision(rev string) *Bootstrap {
	b.version.Revison = rev

	return b
}

func (b *Bootstrap) LogPrefix(prefix string) *Bootstrap {
	b.logging = LogAdapter

	log.SetPrefix(prefix)

	return b
}

func (b *Bootstrap) LogOutput(w io.Writer) *Bootstrap {
	b.logging = LogAdapter

	log.SetOutput(w)

	return b
}

func (b *Bootstrap) GLog() *Bootstrap {
	b.logging = GLogAdapter

	return b
}

func (b *Bootstrap) Syslog() *Bootstrap {
	w, err := syslog.New(syslog.LOG_LOCAL0, "")

	if err != nil {
		panic(err)
	}

	b.logging = newSyslogAdapter(w)

	return b
}

func (b *Bootstrap) SyslogDial(uri *url.URL) *Bootstrap {
	w, err := syslog.Dial(uri.Scheme, uri.Host, syslog.LOG_LOCAL0, "")

	if err != nil {
		panic(err)
	}

	b.logging = newSyslogAdapter(w)

	return b
}
