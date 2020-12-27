package logo

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	stdLog "log"
	"strconv"
	"strings"
	"testing"
)

func TestInfo(t *testing.T) {
	stringBuilder := strings.Builder{}
	Standard.Output = &stringBuilder
	const (
		mode           = ".*INFO.*"
		dateTimeMillis = "\\d\\d\\d\\d-\\d\\d-\\d\\d\\s\\d\\d:\\d\\d:\\d\\d.\\d\\d\\d"
		filename       = "logo_test.go"
		method         = "llog.TestInfo"
	)
	Info(mode)
	assert.Regexp(t, mode+`\s+`+dateTimeMillis+`\s`+filename+`:\d+\s`+method+`:\s`+mode, stringBuilder.String())
}

func TestDebug(t *testing.T) {
	stringBuilder := strings.Builder{}
	Standard.Output = &stringBuilder
	const (
		mode           = ".*DEBUG.*"
		dateTimeMillis = "\\d\\d\\d\\d-\\d\\d-\\d\\d\\s\\d\\d:\\d\\d:\\d\\d.\\d\\d\\d"
		filename       = "logo_test.go"
		method         = "llog.TestDebug"
	)
	Debug(mode)
	assert.Regexp(t, mode+`\s+`+dateTimeMillis+`\s`+filename+`:\d+\s`+method+`:\s`+mode, stringBuilder.String())
}

func TestWarn(t *testing.T) {
	stringBuilder := strings.Builder{}
	Standard.Output = &stringBuilder
	const (
		mode           = ".*WARN.*"
		dateTimeMillis = "\\d\\d\\d\\d-\\d\\d-\\d\\d\\s\\d\\d:\\d\\d:\\d\\d.\\d\\d\\d"
		filename       = "logo_test.go"
		method         = "llog.TestWarn"
	)
	Warn(mode)
	assert.Regexp(t, mode+`\s+`+dateTimeMillis+`\s`+filename+`:\d+\s`+method+`:\s`+mode, stringBuilder.String())
}

func TestError(t *testing.T) {
	stringBuilder := strings.Builder{}
	Standard.Output = &stringBuilder
	const (
		mode           = ".*ERROR.*"
		dateTimeMillis = "\\d\\d\\d\\d-\\d\\d-\\d\\d\\s\\d\\d:\\d\\d:\\d\\d.\\d\\d\\d"
		filename       = "logo_test.go"
		method         = "llog.TestError"
	)
	Error(mode)
	assert.Regexp(t, mode+`\s+`+dateTimeMillis+`\s`+filename+`:\d+\s`+method+`:\s`+mode, stringBuilder.String())
}

func TestPrint(t *testing.T) {
	stringBuilder := strings.Builder{}
	Standard.Output = &stringBuilder
	const (
		mode           = "PRINT"
		dateTimeMillis = "\\d\\d\\d\\d-\\d\\d-\\d\\d\\s\\d\\d:\\d\\d:\\d\\d.\\d\\d\\d"
		filename       = "logo_test.go"
		method         = "llog.TestPrint"
	)
	Print(mode)
	assert.Regexp(t, mode+`\s+`+dateTimeMillis+`\s`+filename+`:\d+\s`+method+`:\s`+mode, stringBuilder.String())
}

func TestConfig(t *testing.T) {
	stringBuilder := strings.Builder{}
	Standard.Output = &stringBuilder
	const (
		mode     = "PRINT"
		date     = "\\d\\d\\d\\d-\\d\\d-\\d\\d"
		time     = "\\d\\d:\\d\\d:\\d\\d"
		millis   = ".\\d\\d\\d"
		filename = "logo_test.go"
		method   = "llog.TestPrint"
	)
	Standard.Config(DateFlag)
	Print(mode)
	assert.Regexp(t, mode+"\\s+"+date+":\\s"+mode, stringBuilder.String())
	stringBuilder.Reset()

	Standard.Config(TimeFlag)
	Print(mode)
	assert.Regexp(t, mode+"\\s+"+time+":\\s"+mode, stringBuilder.String())
	stringBuilder.Reset()

	Standard.Config(MillisFlag)
	Print(mode)
	assert.Regexp(t, mode+"\\s+"+millis+":\\s"+mode, stringBuilder.String())
	stringBuilder.Reset()

	Standard.Config(DateFlag | TimeFlag)
	Print(mode)
	assert.Regexp(t, mode+"\\s+"+date+"\\s"+time+":\\s"+mode, stringBuilder.String())
	stringBuilder.Reset()

	Standard.Config(DateFlag | TimeFlag | MillisFlag)
	Print(mode)
	assert.Regexp(t, mode+"\\s+"+date+"\\s"+time+millis+":\\s"+mode, stringBuilder.String())
	stringBuilder.Reset()
}

func TestJson(t *testing.T) {
	var stringBuilder strings.Builder
	Standard.Output = &stringBuilder
	Standard.Config(1<<6 - 1)
	Print("Print")
	var expected = []string{
		`\s+"level":"PRINT"`,
		`\s+"date":"\d\d\d\d-\d\d-\d\d"`,
		`\s+"time":"\d\d:\d\d:\d\d.\d\d\d"`,
		`\s+"msg":"Print"`,
		`\s+"filename":"logo_test.go:\d+"`,
		`\s+"func_name":"logo.TestJson"`,
	}

	assert.Regexp(t, `{\n?`+strings.Join(expected, `,\n?`)+`\n?}`, stringBuilder.String())
}

func BenchmarkInfo(b *testing.B) {
	Standard.Output = ioutil.Discard
	for i := 0; i < b.N; i++ {
		Info(i)
	}
}

func BenchmarkStdLog(b *testing.B) {
	stdLog.SetOutput(ioutil.Discard)
	for i := 0; i < b.N; i++ {
		stdLog.Print(i)
	}
}

func BenchmarkFmtPrint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Fprintln(ioutil.Discard, i)
	}
}

func BenchmarkWrite(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ioutil.Discard.Write([]byte(strconv.Itoa(i)))
	}
}