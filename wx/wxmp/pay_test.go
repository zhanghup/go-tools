package wxmp_test

import (
	"github.com/zhanghup/go-tools/wx/wxmp"
	"testing"
)

func TestSign(t *testing.T) {
	engine = wxmp.NewEngine(&wxmp.Option{
		Appsecret:     "a0e2253fc4b5fd649a3b6a91ec6ee14b",
		Appid:         "wx1eb1fc2b333e11d0",
		Mchid:         "1606344047",
		MchPrivateKey: `-----BEGIN CERTIFICATE REQUEST-----MIICuTCCAaECAQAwdDELMAkGA1UEBhMCQ04xEjAQBgNVBAgMCUd1YW5nRG9uZzERMA8GA1UEBwwIU2hlblpoZW4xGzAZBgNVBAoMEuW+ruS/oeWVhuaIt+ezu+e7nzEMMAoGA1UECwwDMTIzMRMwEQYDVQQDDAoxNjA2MzQ0MDQ3MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA3gtWaE8iTdw1diZLyafO+tVGnVTs535oPuWQnN5+xMjHHr0SzJF7IP+1Z+hvt0hIABtNnbQhpL0FeVdseBT21Ci4snX0E6z98+W88kF04RzfcTdQQpINh6DWjbwK+PJoLCGMpvDtH1zkTYlhdC/1Uhqm+26ABx95j7mMuXqkY4IiQRD42CDnYRjrrhPiZr+BM5E+SBrJBZiDJ0uYEcUfw4uAAcz+B4TVyqYnaKAeJDhSKf3TisLdT6y5+uwPAkfMz5ZKnYom4UpAbpr1CqKPnXnesit/lVWFiWuBv8ZHe89VoQphp5yadz//P9U6jtH3gCyHzDDoZzwrWaz90jMwOQIDAQABoAAwDQYJKoZIhvcNAQELBQADggEBACdiK04ZV/Li5EhO3mxoByNsbM9jvw1QC1Np0AxpbfyP9r/wJskTd9GHE9yzBxbIWFHJ8N88N3M5kPHauXNko+yXHhCpxWrNem6Bk39kRMUnZzZ+Eyj3xU3xKa6ecGWoR5OfBL/wjpmJilqwTVtoRj+gBnvPpZxPaNeXMR+3pcCHjl8+647IDCoPdGq1CC911VGywKqOLUsXv3uPfTpU6x0RX7wkq4ZF2kjNhDOpJosXr5yT9Txw53DMZTPuBBzmeDefiP7vbEIk1A4VTMkQvsq8kgyXQX/vjVPJgBOs3VcF3IxKibZhIc2nhdnZRelrhs1WmklwiewrwpDjljibS5A=-----END CERTIFICATE REQUEST-----`,
		MchSeriesNo:   "",
	})
	engine.Pay(&wxmp.PayOption{})
}
