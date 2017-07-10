package bench

import (
	"testing"

	"github.com/avct/uasurfer"
	mssola "github.com/mssola/user_agent"
	"github.com/woothee/woothee-go"
	xojoc "xojoc.pw/useragent"
)

var benchmarks = []struct {
	name, device string
}{
	{
		name:   "ios",
		device: `Mozilla/5.0 (iPhone; CPU iPhone OS 7_0 like Mac OS X) AppleWebKit/546.10 (KHTML, like Gecko) Version/6.0 Mobile/7E18WD Safari/8536.25`,
	},
	{
		name:   "android",
		device: `Mozilla/5.0 (Linux; Android 6.0; Nexus 5X Build/MDB08L) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/46.0.2490.76 Mobile Safari/537.36`,
	},
	{
		name:   "kddi",
		device: `KDDI-TS3V UP.Browser/6.2_7.2.7.1.K.6.210 (GUI) MMP/2.0`,
	},
	{
		name:   "googlebot",
		device: `Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)`,
	},
	{
		name:   "vivaldi",
		device: `Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/48.0.2564.88 Safari/537.36 Vivaldi/1.0.380.2`,
	},
}

func BenchmarkWoothee(b *testing.B) {
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				woothee.Parse(bm.device)
			}
		})
	}
}

func BenchmarkUAParser(b *testing.B) {
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				uasurfer.Parse(bm.device)
			}
		})
	}
}

func BenchmarkMssola(b *testing.B) {
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				mssola.New(bm.device)
			}
		})
	}
}

func BenchmarkXojoc(b *testing.B) {
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				xojoc.Parse(bm.device)
			}
		})
	}
}
