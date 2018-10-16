/*
   Copyright The containerd Authors.

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package jwe

import (
	"encoding/base64"
	"strings"
)

func b64Dec(str string) string {
	str = strings.Replace(str, " ", "", -1)
	s, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		panic(err)
	}

	return string(s)
}

// Start JWE Keys
var (
	jwePrivKeyPem = []byte(`-----BEGIN RSA PRIVATE KEY-----
MIIEpQIBAAKCAQEAsl+vKscbGU8TeGLoYJAt+i9zudJtBisoeB4C35EqiRgPXslN
6Qu6AP5FTcR4mBYVzM1txmBELMtZHrZFfRTMRCr3y4JHeNyMescaTOKodLC67hTN
2UKwqjVA4Q+LWBRPpxNoqeWOKDPOJcVSkPpKnBf9lP7iPdBI0s2gCuBwkRbggQqY
HZrio/3MtLzkjo65NKRm9/BrexZpIVIIklYgY78mZYij8WF3XfZ0B4FyjoOn7X4A
x9JARSvOkDzlnNZaRz4WSuXDYgfOpsPwo6YdHejEltf1QOY+DV6rlPoQ0szmCzH5
5Xcufi9/33hvMY0sJ4YWUAvBkGmyw2vHotZCzwIDAQABAoIBAAjOuLOACVKCmQ+E
sryx4dNMrIYsYb3AO8tSkAnB/TuvuHKRtgsfzRtncryYSuwXixQFwLne3v7nO4tM
rLm0YTGsfXfLAwRwv28AjcfmGTNJ1rESzedAZ8C/yGhUgCjlN9mkF7Lr5s0NYcxz
pdQKx8xVUuwcecdblXzzMkfXNTe0uEJP1rTzUy5CKNLaY3rj9L+4pefj1zKRZbJK
z/Om1JdvQ3sev2AcaiY5GOVCq+q/mJpQznhCY/0GZvy1eiJ4jBJXTlI3Ifa3m9Oc
q7qH9G9O0+78wVzU9HZbIMnlxrikyWl23ZXB+NxIlDvArk/d4VCZeNkSrtO7zC0/
45sP96ECgYEA7MkboWnzj+CHGfM/7auHNUl9nMMLy07zfIQInu1SMWgaV+9IM3XU
EFmL77vsUuM7o7vU2FFmyFONQ2dB9CRO1OM8gF2dm0Ckc09OpnC0eMt/z7R7RWbg
d5zx70aZJfeePe8Omj8b8iJ6c/WrhmK5hCOIjk5TApQNkNk1oi4kavUCgYEAwNkm
5Jote8aFpMCZ5G6fwGJWwpHJzOqWFmi3rDfwnIeFkJFKDTXzwAtv90r736LGpMtW
0k569Fna+5ENY2SyrBWUC4Ww4l4FLEDzmgFJsau0VE42r7ng9zm2iKUyWg03/Lze
JgIotvViGS1yPoUUv95uEcXznJ58apz+s57QpDMCgYEArRs5gAAdeAoFuwsCqZbE
+kgH9RsC/Fdz2owMYWPOuyAIYlEkz7pMlsdgbptMYiyN5V3kdWDNa5bpp2VN6lbA
6xJVoOLP3jicAVDxhuzOg6ECh67CkDJt2AR9Oxi5zfABV/X1Dv8kRxi9vRjVlSGH
zvrLUn4gYborUMH7W92v8iECgYEApXYrnpyCRd7RL8howcwAmSpG0m4PvRfRaqyy
WrssYMEOYjmmVati1fV6Pa1CamDZGu+0MIFRkXG/J3UPDaaKfoeNHE26tJ6CxbN8
zzgnqJ9v+52X4jITyUrlSFyk1QrebKUH3Yigsknbv0p06Rt58B3CRtGW8Vwx16+Y
ATlUPm0CgYEAtKPvy6+5eCRigT29ejsO9l7hOwrNhxyWJYqUg7RrajQ/SxbPMm6O
hWuV/5Bb4gGpe7lX18nu4dCsNaxlZj4orfeOw7FslZLoV54krgk7PB6sWnlcSyl0
Hua6v6HMIZ66bmHqc7564uyiEWDFXFN+1k/8RNGPRF0spD4J7/gIx90=
-----END RSA PRIVATE KEY-----`)

	jwePrivKeyDer = []byte(b64Dec(`MIIEpQIBAAKCAQEAsl+vKscbGU8TeGLoYJAt+i9zudJtBisoeB4C35EqiRgPXslN6Qu6AP5FTcR4    mBYVzM1txmBELMtZHrZFfRTMRCr3y4JHeNyMescaTOKodLC67hTN2UKwqjVA4Q+LWBRPpxNoqeWO    KDPOJcVSkPpKnBf9lP7iPdBI0s2gCuBwkRbggQqYHZrio/3MtLzkjo65NKRm9/BrexZpIVIIklYg
    Y78mZYij8WF3XfZ0B4FyjoOn7X4Ax9JARSvOkDzlnNZaRz4WSuXDYgfOpsPwo6YdHejEltf1QOY+
    DV6rlPoQ0szmCzH55Xcufi9/33hvMY0sJ4YWUAvBkGmyw2vHotZCzwIDAQABAoIBAAjOuLOACVKC
    mQ+Esryx4dNMrIYsYb3AO8tSkAnB/TuvuHKRtgsfzRtncryYSuwXixQFwLne3v7nO4tMrLm0YTGs
    fXfLAwRwv28AjcfmGTNJ1rESzedAZ8C/yGhUgCjlN9mkF7Lr5s0NYcxzpdQKx8xVUuwcecdblXzz
    MkfXNTe0uEJP1rTzUy5CKNLaY3rj9L+4pefj1zKRZbJKz/Om1JdvQ3sev2AcaiY5GOVCq+q/mJpQ
    znhCY/0GZvy1eiJ4jBJXTlI3Ifa3m9Ocq7qH9G9O0+78wVzU9HZbIMnlxrikyWl23ZXB+NxIlDvA
    rk/d4VCZeNkSrtO7zC0/45sP96ECgYEA7MkboWnzj+CHGfM/7auHNUl9nMMLy07zfIQInu1SMWga
    V+9IM3XUEFmL77vsUuM7o7vU2FFmyFONQ2dB9CRO1OM8gF2dm0Ckc09OpnC0eMt/z7R7RWbgd5zx
    70aZJfeePe8Omj8b8iJ6c/WrhmK5hCOIjk5TApQNkNk1oi4kavUCgYEAwNkm5Jote8aFpMCZ5G6f
    wGJWwpHJzOqWFmi3rDfwnIeFkJFKDTXzwAtv90r736LGpMtW0k569Fna+5ENY2SyrBWUC4Ww4l4F
    LEDzmgFJsau0VE42r7ng9zm2iKUyWg03/LzeJgIotvViGS1yPoUUv95uEcXznJ58apz+s57QpDMC
    gYEArRs5gAAdeAoFuwsCqZbE+kgH9RsC/Fdz2owMYWPOuyAIYlEkz7pMlsdgbptMYiyN5V3kdWDN
    a5bpp2VN6lbA6xJVoOLP3jicAVDxhuzOg6ECh67CkDJt2AR9Oxi5zfABV/X1Dv8kRxi9vRjVlSGH
    zvrLUn4gYborUMH7W92v8iECgYEApXYrnpyCRd7RL8howcwAmSpG0m4PvRfRaqyyWrssYMEOYjmm
    Vati1fV6Pa1CamDZGu+0MIFRkXG/J3UPDaaKfoeNHE26tJ6CxbN8zzgnqJ9v+52X4jITyUrlSFyk
    1QrebKUH3Yigsknbv0p06Rt58B3CRtGW8Vwx16+YATlUPm0CgYEAtKPvy6+5eCRigT29ejsO9l7h
    OwrNhxyWJYqUg7RrajQ/SxbPMm6OhWuV/5Bb4gGpe7lX18nu4dCsNaxlZj4orfeOw7FslZLoV54k
    rgk7PB6sWnlcSyl0Hua6v6HMIZ66bmHqc7564uyiEWDFXFN+1k/8RNGPRF0spD4J7/gIx90=`))

	jwePubKeyPem = []byte(`-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAsl+vKscbGU8TeGLoYJAt
+i9zudJtBisoeB4C35EqiRgPXslN6Qu6AP5FTcR4mBYVzM1txmBELMtZHrZFfRTM
RCr3y4JHeNyMescaTOKodLC67hTN2UKwqjVA4Q+LWBRPpxNoqeWOKDPOJcVSkPpK
nBf9lP7iPdBI0s2gCuBwkRbggQqYHZrio/3MtLzkjo65NKRm9/BrexZpIVIIklYg
Y78mZYij8WF3XfZ0B4FyjoOn7X4Ax9JARSvOkDzlnNZaRz4WSuXDYgfOpsPwo6Yd
HejEltf1QOY+DV6rlPoQ0szmCzH55Xcufi9/33hvMY0sJ4YWUAvBkGmyw2vHotZC
zwIDAQAB
-----END PUBLIC KEY-----`)

	jwePubKeyDer = []byte(b64Dec(`MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAsl+vKscbGU8TeGLoYJAt+i9zudJtBiso
    eB4C35EqiRgPXslN6Qu6AP5FTcR4mBYVzM1txmBELMtZHrZFfRTMRCr3y4JHeNyMescaTOKodLC6
    7hTN2UKwqjVA4Q+LWBRPpxNoqeWOKDPOJcVSkPpKnBf9lP7iPdBI0s2gCuBwkRbggQqYHZrio/3M
    tLzkjo65NKRm9/BrexZpIVIIklYgY78mZYij8WF3XfZ0B4FyjoOn7X4Ax9JARSvOkDzlnNZaRz4W
    SuXDYgfOpsPwo6YdHejEltf1QOY+DV6rlPoQ0szmCzH55Xcufi9/33hvMY0sJ4YWUAvBkGmyw2vH
    otZCzwIDAQAB`))

	jwePrivKey2Pem = []byte(`-----BEGIN RSA PRIVATE KEY-----
MIIEpQIBAAKCAQEA515unjMo31umXoIjTsKqNyrKYv3DVrNQQqzCHDnVYuhxrU81
6Gj/sW03DBoqYf4Ic7ohbcJFiCzgLU4tYeOT7GZpBZT1UIulHUCg1m/pM3zEpD1U
VRmerfLBuuK4osAF/f1uYddFT/2/p9BiHCxExGcd8bN+Qe2hh7Uz/Tw0tUFNtRFh
lc/MipQhzwHXSr8y1mW5YDNh5Zk4Uz+fyUEnRPLS+MQXn82IrId+gUbV+2uptHyj
lrLpwcxYmk8aB1EQXJevBcFhWI95bIfsm117Y7Oy4OLR/3LBXN5nN0CZfAhOBwBR
vYaJ0RRNFWI/AXHjd32Ron8j16lIDKdS9I9qlwIDAQABAoIBACpthLd6Bjq/Ycje
8H6W8APh0u9IPbP+ee9gItBuQpU6ru3rIqWV652ru3Z6rd7+aKpgLZUlFP8dy5ZD
ScszooKtXQDrCflVQlgU6+mm6ArLDHxZysc4RYL8i04sGVOvBupGuSE0Cr/adnTd
n8Au1gV2K2WRVsvcOczbC8eabMf38mvtsAYn6/Ehfd63IZ0JjurCblK6d477E8V9
iu0XiOLHE9sFHMGkWg2K1QozZDXjVjRMdInVcbc/NZGMgpAj1LHHgUZkr5yWqOMl
IP5R0R8z/6LH/QeXtEQYqORr5K4ONEelH0HYonHCdHjtz12fkErFyAX8mh/25zoV
TJn0JoECgYEA+0bxA8vpX1dmtylRj3GlBEnQFWjTOwCwh5lckcmG8NKBtG752nuz
xSD5fCIsKqzsZAyoItCZ3QtA0K+ryWkYl/gL2DYSPt/rvyw8mEhvBQTbFbfBR9Sy
PeyOjIWcqADvu99w4JcngjB+lR50PWrfkIbwdybzWViMbI/2LZ2fW3cCgYEA67ey
9qaVZN61TE9E1FkvDuN3KA2DH2ho/mvlWKsF9ft+6LqXlsLbhxgVFxjX0RFqGLeq
DYb1/0dKHknEq6V0ZQRNZ1JuzeuTduZ8vuKdBI6JkWK86oAykGTq0nLTw0IINE3J
qSbifugYgRAPL/6hXVPbTh9BUam8duxthUwD8eECgYEAmLt/JbqdCGmcsno37AO8
tMWU6F6F/hgmNNXAEZE4J0scsaq+zdFg7NJlMtGmnO3s5cdXr4mx7Ey5wd71gQAT
hdOsh2geYP9EUTg3QKzOZnOUIzhFED81dDREVR+ln+ypyz0+ZBUcW6LUXhlbuDUs
3LFYmmQfiFAtUpOSpBlp0nMCgYEA0nGI12hWDF5AokZK/wI4XyR5J0sY+5tt0Wdm
tMjLY5cK8KBV4gVJlMzNV3eYhlDz1elzauxJB4YQCAZ4DX6D8gPrTwlrX3CokQip
6onLAVx4OVJbs0iM3BkdBJH7uWFkjb29AsVyhTaVWfSKeqDsU7QgIRkKaewOFGZ0
SQNaqkECgYEA94y65SrkydMBcqlZ1+qxNJcFlLTZ7abkuZb7C/Tyb9gRVzlLujfD
ZZ5a0Bv28a5z1iUa9TDu7gc+3hBq4quE2i8BZhAEVYsrD4HhhEeu2qpk8u5M05il
4WAAWmgaMLm6abjz2sbcvVb9GAaM/M5ZE6G4cetNQhrGJBa3FS12zlI=
-----END RSA PRIVATE KEY-----`)

	jwePubKey2Pem = []byte(`-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA515unjMo31umXoIjTsKq
NyrKYv3DVrNQQqzCHDnVYuhxrU816Gj/sW03DBoqYf4Ic7ohbcJFiCzgLU4tYeOT
7GZpBZT1UIulHUCg1m/pM3zEpD1UVRmerfLBuuK4osAF/f1uYddFT/2/p9BiHCxE
xGcd8bN+Qe2hh7Uz/Tw0tUFNtRFhlc/MipQhzwHXSr8y1mW5YDNh5Zk4Uz+fyUEn
RPLS+MQXn82IrId+gUbV+2uptHyjlrLpwcxYmk8aB1EQXJevBcFhWI95bIfsm117
Y7Oy4OLR/3LBXN5nN0CZfAhOBwBRvYaJ0RRNFWI/AXHjd32Ron8j16lIDKdS9I9q
lwIDAQAB
-----END PUBLIC KEY-----`)
)
