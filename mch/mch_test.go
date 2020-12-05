package mch

import (
	"testing"

	"github.com/shenghui0779/gochat/wx"
	"github.com/stretchr/testify/assert"
)

func TestLoadCertFromPemBlock(t *testing.T) {
	mch := New("wx2421b1c4370ec43b", "10000100", "192006250b4c09247ec02edce69f6a2d")

	assert.Nil(t, mch.LoadCertFromPemBlock(certPemBlock, keyPemBlock))
}

// 涉及时间戳，签名会变化（已通过固定时间戳验证）
// func TestAPPAPI(t *testing.T) {
// 	mch := New("wx2421b1c4370ec43b", "10000100", "192006250b4c09247ec02edce69f6a2d")

// 	mch.nonce = func(size int) string {
// 		return "5K8264ILTKCH16CQ2502SI8ZNMTM67VS"
// 	}

// 	m := mch.APPAPI("WX1217752501201407033233368018")

// 	assert.Equal(t, wx.WXML{
// 		"appid":     "wx2421b1c4370ec43b",
// 		"partnerid": "10000100",
// 		"prepayid":  "WX1217752501201407033233368018",
// 		"package":   "Sign=WXPay",
// 		"noncestr":  "5K8264ILTKCH16CQ2502SI8ZNMTM67VS",
// 		"timestamp": "1414561699",
// 		"sign":      "C9612FA7A6BA5F51E195D5F9337CA288",
// 	}, m)
// }

// 涉及时间戳，签名会变化（已通过固定时间戳验证）
// func TestJSAPI(t *testing.T) {
// 	mch := New("wx2421b1c4370ec43b", "10000100", "192006250b4c09247ec02edce69f6a2d")

// 	mch.nonce = func(size int) string {
// 		return "e61463f8efa94090b1f366cccfbbb444"
// 	}

// 	m := mch.JSAPI("u802345jgfjsdfgsdg888")

// 	assert.Equal(t, wx.WXML{
// 		"appId":     "wx2421b1c4370ec43b",
// 		"timeStamp": "1414561699",
// 		"nonceStr":  "e61463f8efa94090b1f366cccfbbb444",
// 		"package":   "prepay_id=u802345jgfjsdfgsdg888",
// 		"signType":  "MD5",
// 		"paySign":   "A62A01211E36F5D2173A9EE93EBAC56C",
// 	}, m)
// }

// 涉及时间戳，签名会变化（已通过固定时间戳验证）
// func TestMinipRedpackJSAPI(t *testing.T) {
// 	mch := New("wx2421b1c4370ec43b", "10000100", "192006250b4c09247ec02edce69f6a2d")

// 	mch.nonce = func(size int) string {
// 		return "e61463f8efa94090b1f366cccfbbb444"
// 	}

// 	m := mch.MinipRedpackJSAPI("sendid=242e8abd163d300019b2cae74ba8e8c06e3f0e51ab84d16b3c80decd22a5b672&ver=8&sign=4110d649a5aef52dd6b95654ddf91ca7d5411ac159ace4e1a766b7d3967a1c3dfe1d256811445a4abda2d9cfa4a9b377a829258bd00d90313c6c346f2349fe5d&mchid=11475856&appid=wxd27ebc41b85ce36d")

// 	assert.Equal(t, wx.WXML{
// 		"timeStamp": "1414561699",
// 		"nonceStr":  "e61463f8efa94090b1f366cccfbbb444",
// 		"package":   "sendid%3D242e8abd163d300019b2cae74ba8e8c06e3f0e51ab84d16b3c80decd22a5b672%26ver%3D8%26sign%3D4110d649a5aef52dd6b95654ddf91ca7d5411ac159ace4e1a766b7d3967a1c3dfe1d256811445a4abda2d9cfa4a9b377a829258bd00d90313c6c346f2349fe5d%26mchid%3D11475856%26appid%3Dwxd27ebc41b85ce36d",
// 		"signType":  "MD5",
// 		"paySign":   "0cecd02326e26c27fbc77f6062ef8654",
// 	}, m)
// }

func TestSignWithMD5(t *testing.T) {
	mch := New("wx2421b1c4370ec43b", "10000100", "192006250b4c09247ec02edce69f6a2d")

	m := wx.WXML{
		"appid":     "wx2421b1c4370ec43b",
		"partnerid": "10000100",
		"prepayid":  "WX1217752501201407033233368018",
		"package":   "Sign=WXPay",
		"noncestr":  "5K8264ILTKCH16CQ2502SI8ZNMTM67VS",
		"timestamp": "1514363815",
	}

	// 签名校验来自：[微信支付接口签名校验工具](https://pay.weixin.qq.com/wiki/doc/api/app/app.php?chapter=20_1)
	assert.Equal(t, "66724B3332E124BFC3D62A31A68F7887", mch.SignWithMD5(m, true))
}

func TestSignWithHMacSHA256(t *testing.T) {
	mch := New("wx2421b1c4370ec43b", "10000100", "192006250b4c09247ec02edce69f6a2d")

	m := wx.WXML{
		"appid":     "wx2421b1c4370ec43b",
		"partnerid": "10000100",
		"prepayid":  "WX1217752501201407033233368018",
		"package":   "Sign=WXPay",
		"noncestr":  "5K8264ILTKCH16CQ2502SI8ZNMTM67VS",
		"timestamp": "1514363815",
	}

	// 签名校验来自：[微信支付接口签名校验工具](https://pay.weixin.qq.com/wiki/doc/api/app/app.php?chapter=20_1)
	assert.Equal(t, "3B12F569A5714858F8251366BC3CBCDDBD249905CCA01D8F56D365EF1FC2CA5C", mch.SignWithHMacSHA256(m, true))
}

func TestVerifyWXMLResult(t *testing.T) {
	mch := New("wx2421b1c4370ec43b", "10000100", "192006250b4c09247ec02edce69f6a2d")

	m := wx.WXML{
		"return_code": "SUCCESS",
		"return_msg":  "OK",
		"appid":       "wx2421b1c4370ec43b",
		"mch_id":      "10000100",
		"nonce_str":   "IITRi8Iabbblz1Jc",
		"sign":        "E515C9BE3D3129764915407267CA0243",
		"result_code": "SUCCESS",
		"prepay_id":   "wx201411101639507cbf6ffd8b0779950874",
		"trade_type":  "APP",
	}

	assert.Nil(t, mch.VerifyWXMLResult(m))
}

var (
	// tls certificate
	certPemBlock []byte
	keyPemBlock  []byte
	// rsa key
	privateKey []byte
	publicKey  []byte
)

func TestMain(m *testing.M) {
	certPemBlock = []byte(`-----BEGIN CERTIFICATE-----
MIIEazCCA9SgAwIBAgIDHEZcMA0GCSqGSIb3DQEBBQUAMIGKMQswCQYDVQQGEwJD
TjESMBAGA1UECBMJR3Vhbmdkb25nMREwDwYDVQQHEwhTaGVuemhlbjEQMA4GA1UE
ChMHVGVuY2VudDEMMAoGA1UECxMDV1hHMRMwEQYDVQQDEwpNbXBheW1jaENBMR8w
HQYJKoZIhvcNAQkBFhBtbXBheW1jaEB0ZW5jZW50MB4XDTE2MDMyNTEwMjAwMloX
DTI2MDMyMzEwMjAwMlowgZsxCzAJBgNVBAYTAkNOMRIwEAYDVQQIEwlHdWFuZ2Rv
bmcxETAPBgNVBAcTCFNoZW56aGVuMRAwDgYDVQQKEwdUZW5jZW50MQ4wDAYDVQQL
EwVNTVBheTEwMC4GA1UEAxQn5Y2X5Lqs6JOd6bK45Lq6572R57uc56eR5oqA5pyJ
6ZmQ5YWs5Y+4MREwDwYDVQQEEwgxMTQ1OTgyNTCCASIwDQYJKoZIhvcNAQEBBQAD
ggEPADCCAQoCggEBALROHwnq98ftW1tsfi1ymsav+bAa2/Wq6oNuPXNCuHRwcpXB
KCQa5iThh64Ud9UnO87fzZ2WHD9sacXAtbdh5m9IfYMXGIQMzHIkTyix94paFO6v
wFJFkEJlwKJyg3AymXTDB/cNWXhZL/idz+ymy0wnuGuW1IVt0fa6eVQK1E7WNDi6
dEG0GEX1NnxeEEoP6Pa+XGT3g+zgI5G0diRTTlDKiJhKgl+589JE6AFe6JqiVdIc
5bzoaSzWdCkD7JfwvmRggbXRSsAQ2QMouqaeIMpwr5axkvEybleu2+mReqVB5pwE
0+TwF56fbiAZCkc9y16qxleDRHsw3krGU/qb0wECAwEAAaOCAUYwggFCMAkGA1Ud
EwQCMAAwLAYJYIZIAYb4QgENBB8WHSJDRVMtQ0EgR2VuZXJhdGUgQ2VydGlmaWNh
dGUiMB0GA1UdDgQWBBRuLat+HKTimCUM74piXoMKLWPRTDCBvwYDVR0jBIG3MIG0
gBQ+BSb2ImK0FVuIzWR+sNRip+WGdKGBkKSBjTCBijELMAkGA1UEBhMCQ04xEjAQ
BgNVBAgTCUd1YW5nZG9uZzERMA8GA1UEBxMIU2hlbnpoZW4xEDAOBgNVBAoTB1Rl
bmNlbnQxDDAKBgNVBAsTA1dYRzETMBEGA1UEAxMKTW1wYXltY2hDQTEfMB0GCSqG
SIb3DQEJARYQbW1wYXltY2hAdGVuY2VudIIJALtUlyu8AOhXMA4GA1UdDwEB/wQE
AwIGwDAWBgNVHSUBAf8EDDAKBggrBgEFBQcDAjANBgkqhkiG9w0BAQUFAAOBgQB8
JjIjIbrLTsKeyhtUwosT26vAQlyIdZVKaX7iHmt6HKjZKpi6qziIMFYWj/K2AutE
WGYW0ex09v5KORVBi4ahyJnDFyPC6k/5Dhe++4y4SPxJ/2EI7b0mpPxAF16VePt+
2RhogAbMS+gv7ecrPv/H1jU+lvZR3ygxHnaG3BP3PA==
-----END CERTIFICATE-----`)

	keyPemBlock = []byte(`-----BEGIN PRIVATE KEY-----
MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQC0Th8J6vfH7Vtb
bH4tcprGr/mwGtv1quqDbj1zQrh0cHKVwSgkGuYk4YeuFHfVJzvO382dlhw/bGnF
wLW3YeZvSH2DFxiEDMxyJE8osfeKWhTur8BSRZBCZcCicoNwMpl0wwf3DVl4WS/4
nc/spstMJ7hrltSFbdH2unlUCtRO1jQ4unRBtBhF9TZ8XhBKD+j2vlxk94Ps4COR
tHYkU05QyoiYSoJfufPSROgBXuiaolXSHOW86Gks1nQpA+yX8L5kYIG10UrAENkD
KLqmniDKcK+WsZLxMm5XrtvpkXqlQeacBNPk8Been24gGQpHPcteqsZXg0R7MN5K
xlP6m9MBAgMBAAECggEBAKacyfHQPsdwfkstJiu5C20uj/w71aZeGfb5l686qFhw
0HGx2/YBJUpPXaFvKIy/hHTWOpq0a8Xv2I30VfbvcJDE27aXUQA1E3cmNj/UtHoU
Y+NsZLuhrHyuqiNyziKPn15WGrYgj9y2Da4fplN4jcQBsFk7N4dUxADKr/MJTsbJ
ewkre7WLHaooylpB0ILi9W1cBsjG+z4xNaJcgye/2GlxWjRI2EUqXzSPXRpT0F2M
pzg5s8JyPYXxKDktMBvyKLejFz61ULRYER0wC3/1Xpx5yWkDL1VbGvfRHEctFf4k
ISAE6MFwlnKfBuW9PvEvo1xy7XgK7PClPI1LxXXgokECgYEA6/QwSpjqwxZ2WFRH
IuFRYlTdfIgLuQ2nXKETNEfiYurK8WSD8hhSIf7YI6woUFaDk4a+ID5VONI48UwR
MenYiKop6SgO3DQiOr3WzLX/n3/nGp5WCwYRlhZ8luuF2ro3vXwb7cBjz0Q/IbN6
cAhrQWuFOgpj9/oM5Q7v2nkQNA8CgYEAw5+d9ZNig20EOcopAhXxn4VUg3Pkido7
VMGdPlN3er8Ib3n34mWPTNvz9ouFRvZrb/YlTGAuaYn3NBd/XyF4khNxLEMJc5QU
M1EmvciTSJraUE7x/AMpkX7lfGsaCwFRC6+KtqLIpc/EoRVrhJvTvLFg49e9Ripa
58XqAnX0N+8CgYEAtbrBZvMX/WHTjHx8vZSkxgNvA3cU8FZfzKwSynWDG4STAhDU
vyWUPLK3beIuupGnjXx9+v+HS9g+GzrnE3Z0W+4TkYxUxa0xn+SPB+Q1GXe9W3cP
9jWaXeq70tFbqvc57ysjp55CQWTi6uX9K0SQtzZEyksua9OfEtzKR45uuGkCgYBu
9j0tLRq2HcJF3WwBaN0TdReJMNWzc/wviVteHQ4qq+1m/jIjUyRDnof1kxJYPDKY
4XAqsGvjJYT5IVL4bQ3tHeYWCzhzPM1whlmJURpqFpw67WzJXCnaA7a8KiwtjeOg
00PFcPSLSRzpmjLQl2s1HxAsbRVJlYDW8yZXmvyaNQKBgHEPplp8BNBN3mKPeZ4m
+ruMNs7DbFCA0b+QADEqSabD2xH25lf3LdB/xB0+CJZlBE6sxBbOE/FKkIST8ZUq
TeF56obd7Ld7cYJEkNoqjyk8fH5ZLtT3AXlgIPc1zCmB9IhFFaUVsqqaAWfoBwA4
fyBMdIgBndcj4ZDklE0z68SP
-----END PRIVATE KEY-----`)

	privateKey = []byte(`-----BEGIN RSA PRIVATE KEY-----
MIIEogIBAAKCAQEAl1c+37GJSFSqbuHJ/wgeLzxLp7C2GYrjzVAnEF3xgjJVTltk
Qzdu3u+fcB3c/dgHX/Zdv5fqVoOqvoOMk4N4zdGeaxN+Cm19c1gsxigNJDtm6Qno
1s1T/qPph/zRArylM0N9Z3vWVEq4xI4B4NXk6IoK/bXc1dwQe5UBzIZyzU5aWfqm
TQilWEs7mqro43LTFkhN05QjC7IUFvWEhh6TwvGYLBSAn+oNw/uSAu6B3c6dh+ps
lgORCzrIRs68GWsARGZkI/lmOJWEgzQ9KC7byHVqEnDDaWQFyQpq30JdP6YTXR/x
lKyo8f1DingoSDXAhKMGRKaT4oIFkE6OA3jtDQIDAQABAoIBAEYkE2nNeJjjVJJL
Zzkh5At0YIP5rIwuCOJRMamuQI5dEZhdVxyoBmSgnj5yOMgVZWWeY1i26c7q+ymN
AowjtMt/SXLK9/GRSUE6LpYwXdbbCRkclKNpSnKMOWHjVGN2VwJpVyetB2rNrtC6
GDjCKXN09x8bOJyNf74nE0xdS7vGzDlmAhCwju34DuMhdj8GBtLZo8O0esaeqNuK
EhlQrur9KuyYJR63ZR306qJpVE7ZX6bFQZpwTrebnATHDnWcvVbVWWpfe8xmQwNa
b2Gsctv8Ght/Ka/OjbRP0d48ZnTGeOuC9eKjpUKi2nZiEiYsCUjTxO30Ib6Pw2Z3
lWMx7kECgYEAxM2UtYjTXFcIbRWSGx9b997xhPpnxLSPzO4JIM2WdQqlRBdgOi7u
BNIL19Z37d6CElEYJ+G/6lqs072xMWt4Nph2cgiKUzcOAAKfS0vna/IXir4oGhTb
auAsj7Ga7dQi23a3UTDb1bNavemo3SqYI1anud00TnyQdBvVJ1ZwADUCgYEAxNzv
zDLiABRETLtFU7zOEjYsB/+WV2cvofsqvq8NQDdyOP6UVZ8vE/DkG61uyMpWp0u/
3/A9krLTz9Gfgw4A7CFFDV3S+z1AY1T2N7I04+QQHMqfbcjotVEG7xouuEfjDN2P
Xi5M2zcmTAkuStO7Yx5UdGPdJNv6JgJyy2doBHkCgYAu6i8kI2z3W0wH7Rd6Xbxn
137Ny3/HNZ/+I1SLvFa8qgABvmzTEfLttUDbgCXwz5VEVo6imz9L17fRdivycwMi
SLAbuQt4kOxGdlmQ8pRFeF3CVlhq90PjM3OMAbPENEjm9mL2+OW/CNV95mC58Hh6
HCM5vJDGkQ1CkIv8p69lbQKBgAYRWULN/rFJ7qD+1LA0DZX6HXlRo2ymPY2clEC0
XJAyJU8kaaYJ9gWDU0SXH+cIdYtKhmt8mClBYc3yBByh/d1JWTuEPNCJnsZxA/XL
hF3R1b1NcYSMwL918+TCxdXgQVtQKO8aNjw7gu6tCcQ8qnXvpWLBATv1m8w4Hxmt
4kLhAoGAejdp4xTh6OYb4kfZA5EN/9wBO3l/7TwWrOe8qT1/FtWMfmcU62Y3LdXE
xuHKcd+Q3/PUQKM5lPFpXqyY/pCE9AQpjFmjo5eU99NNy/oS0P8IaCS2SyppGhF2
HsIxLjl3+jtjS8cptPO47qFnr7Pnvb7kA8MNVrI+ymny/WG/yfU=
-----END RSA PRIVATE KEY-----`)

	publicKey = []byte(`-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAl1c+37GJSFSqbuHJ/wge
LzxLp7C2GYrjzVAnEF3xgjJVTltkQzdu3u+fcB3c/dgHX/Zdv5fqVoOqvoOMk4N4
zdGeaxN+Cm19c1gsxigNJDtm6Qno1s1T/qPph/zRArylM0N9Z3vWVEq4xI4B4NXk
6IoK/bXc1dwQe5UBzIZyzU5aWfqmTQilWEs7mqro43LTFkhN05QjC7IUFvWEhh6T
wvGYLBSAn+oNw/uSAu6B3c6dh+pslgORCzrIRs68GWsARGZkI/lmOJWEgzQ9KC7b
yHVqEnDDaWQFyQpq30JdP6YTXR/xlKyo8f1DingoSDXAhKMGRKaT4oIFkE6OA3jt
DQIDAQAB
-----END PUBLIC KEY-----`)

	m.Run()
}
