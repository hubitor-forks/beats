// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package kubernetes

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "kubernetes", asset.ModuleFieldsPri, AssetKubernetes); err != nil {
		panic(err)
	}
}

// AssetKubernetes returns asset data.
// This is the base64 encoded gzipped contents of ../metricbeat/module/kubernetes.
func AssetKubernetes() string {
	return "eJzsXV9v4ziSf+9PQfRT5pD1w+FwD3k4YCeziw26pzeXdO88HA4eWirbnMikmqSSeD79gtQ/WiIpyqKddCI9NDq2VfVjFf9UFYvFv6AH2F+hh2IFnIIE8QEhSWQGV+jjp+bDjx8QSkEknOSSMHqF/ucDQgi1P0A7kJwk6m0OGWABV2iDPyAkQEpCN+IK/d9HIbKPl+jjVsr84/+r77aMy2XC6JpsrtAaZwI+ILQmkKXiSjP4C6J4Bx146pH7XHHgrMirTyzw1HND14zvsPoYYZoiIbEkQpJEILZGOUsF2mGKN5Ci1d7gs6gomGhMRDgnAvgj8OYbGygPsI78/np7g0qChijr51Ck9dOFZsLj8L0AIRdJRoDKg5/UOB9g/8R42vnOg1Y915oegmdICqXXmpHwouAgWMETiIfjrqQMKbLS7gIQxeqUGFzkezASlscHgDRZdJFkhZDALzVTkeMELhvp/OTF9Qh8FQ/WP75+vUU9kr2eydKIotA8eyT7PKkEKpeKUXw1VBg0C9Rj0cWS8v2SFzQejN9AboEjuYWaByoECJTyPeoy6oJ5ILTLbQKST4SmanatqA+oZJczGneOqkmiLaZppmYpQyheNN25eyISNalrkmjNas0ETBOPwAVhEbtGRbBB0W9mF4KW3MHiNhFCPUhshLvMdyC3LGJ/1APTQrTXaCYidsOmxV2qNducswSEsHK0dUTbem/SS/JiISDpfV/TTFmxyrrzXq8h17ffkICE0bSLrOW0gx3je7WskxSoXKz2rWXW55sxurF8WdplV8j18gGqn9WPEKGo5llhGIL4SLgscHZOhBXLIYDrVCxYDnSRsKI3+w1CO2D9pditgKsZVxFEa5JB8wPG3WoUEnMJaYROc192GCQITUBPMVXnrnlYB4ByBKL1/mZdLbi29heFWOTAE6CSZLD4D2cL2eoPSGwKKL9YjpFDPeZrEGhHEs6q4YRaOG6d2Johit1E/fhxJcWuyLAkj4BsrHzQpnfeGpqmpFeomv4gEEH+hHJkx9T0GNAKwSi1GpB9Wo0xIR1gHKliA+YpNKzIezCInFEBL6reEsIY/fZBn17BJspgDfeBxlBxBcVOqm/0x+9TdcOsK00ZBln4+Dt5O5baOvCBsECWKEunyfGMvIjWgjV2YzLLsASa7I/pyTZtiZrgpeqiCkH5NykNJ3NNGoQUrws1mOh4wayK5AHkWZecijXaEiHZhuMdKkG4wYaaEmNQ1DRLTYYq7zSWQ4uFmoZw+WEYmBfQY4s6XJNJwbmax6bL7oauM7LZyoCuzuiGF5QSuonqqrTzZ6IXLfU2qhj5o8ogk3RRyj3KTN4G/SttCoSl5mJlj4uUyAU8uhQxlr2mhzQ9e3tLhhwUNEgj8qxJdpm3aw2VmNBpexyGdBt6UbY4tGe5lGRnD+WmWHa/GAjY3CuCqEfQCK8Er+JDEcrbb6gQeAMWQbiabULR7zrHoQ2Qj+pBIxm3ER4mPsTAZGKZlLtsHHNJ/QzI13yum06npH7NOFSip5g6F6wDtJgyJRYX6EHAgWDLTgHpAMMGFkthkVvXpBaVSHAG6XKdMez6Ye1yVF5OjDYo6WKBcE1T/c3WOiwkmcSZxo5wlrEES7zKQL3nbWxGdkT+eK1NYU0opCX8JvreToMX6hOnRBBZo4LqdyG1b+BlbBMePx5o1We2UWb4mo2cjPAjJhm2B6GmT0guTxiFjLwhdxqF61pLp2kqSnCOEyL3yvS1U29m1OqXb186ZU8Ol4ya7N6+VPSUHi4UomYC907FtLXdbr2jiIvYV90H2nHibI6xEcLBb3LEQqUYhQBy9Mv4gHTXsAA63MOKFjp6HxN1twcObMOdzpR+XQIpxeBs7iu3K3810I80LR36R6/eugxp8wQDs+oQbhvTlBDvpSmgNzVG7u7v/SOkBvzE+AOhGwHuMNhbkMdvZTORABkmlxxvYI2LzBJIHBMetCNq41aKDXLwaVZN/AfjZ8KjeTlRNaOHMbmOmOfzHjyKO8akzmQReyFhN9q5eB/Gjl1Kpvn93n0wu4Qqy/vlfLEz+BjfLN6FGdnnLMuAl4cfJkX4rxti1VEKb3x/BTI0wv8iSajnzEs/d6LrmRNc1b/x2H3BOwjLo/6T0Yh8b+iaYyF5kciCQ5/4nM5bNmdO553Teed03oBmzOm8diBzOm8wxjmdd07nndN5p6fzWqzMsQm+T4w/fC+gsFucxyx9CjQog7NMupu+nH8uCTbZddVi7rMlCromlIhtFHPiW0MshDVO0xh9+LdaL4rgQEdOIZfbqDw1xcHhIzmJMl5bvmYOs6Zud8xYCotEueyJZHb/+piOC48k0ZZETBtYb1zUlH0ddgs4k9sYmeEt84YqsoeCTpGV7+dU4nFsVoWzuz3YSnI3spmTAKfAF0Qsd1hIR0xmxVgGuGvoDR1b37bn1rWuiUAdHh+6aHS+6ocu+xEhq69bMItvlPmvddQK1Dqkx0bzjdxiiTAHtAEKHMuyWkidLVzNqwccCFWOrRLup27tEjQi3dXdwRy69kr7ulxeFRfEIWE8FaXcm84nyQ7Kz3LMJUmKDPNSCGiLBWKJTkFPLQj1mxLvcgvK/mTiC/utCRdyWbGijood49N7v9YAVTs1D9TyUJ91e5V53OPkgBSLATxtLET09uJKDBKeZXhv+LWkU/UESNvyAOQRqEUcCcv3S8lsCNo1DYuOq+cOvXnR3WlKoeCaXtgtu3Ek96/7vNlk93PcgcQpPohpu7v9gD5KSggLwRKiZ5knIrdenfgGkn1Ijl/hm0mIA+4Gf5BvAATsVRwMAs2AMOqX/EkDzBVnP09dXCcuY00SEYqetiTZVlPuExbtimNFU8fBl9FrhvyrqhliCsQfdi9IxK2Mb5R8LwDp4DBZE2UgMAOIJTjQhEEhWy8zQh8igrn7jDjkHIRCU9WTcU0IhD6y7BHSpQXjqeaFmqdNLr4ZAuckfs/56+1NU3Gm6j0edcUtPaR4P1TlhwYYx508qDF5eJiebrzWlEeIPu6A/XbzywBv0/2cYr0bR8q0xzCfJptPkzme2KfJvqj+9mMfJJuzy23PnF3eeeJll89JxB3AcxKxHficROxJIqYgVb+JNl/z5zfd+e4gAfKo47QuWk00mXPbflQg5lA8zy4+TbTmbSvkK8dU7IiUr0cnX606acLQc8Z++QRK8+9zsv5IAc15+u3TE857SNE3NpodR4G7oM5xhrtF9TpOb7d4XCe4G5umoM4IzjHzNtkpC/BEp/Hda8IwgyEmKHCEo9AQSchIR+NCKTc7bfGOXzVQ4MqB3rMYA9YWNGaye4citK9AjbN6cMZmSgw7Z+kPGcKePdLymT3S9vmRFPLDeaTvYs/oleyS9GC9xrIoY8rtvasSe2pJbaqgiG4ZlLDaepH3x+atoA7sVzmi5kJD8YbZ0dWG3kc48GC4uJvc2TRcvvVdw1IsT729Q7fb8Ma3lUuBNAfllUT0CcEBseR4A8uT7V6WoIJ3UpfnQOPeRzVKNDzvp3jtxpkRTSvgytbBeifNyR5LUZKjc+dddU7aOHIaJU/eVt/EyIzvFiKZwqVHrhFct/7IVKkd0vPV9xhzvGW4sof3TGVgVY9xNT08A88/gx1TzWNULY/IyLxVPAJreHggTajfEVK9I7xjjKnc4azbcVyvHl2xw3vMP6RaR5RaHWMrdcRC5D3gP75GR2jnDK7PcWx1jnCthoMdqNswsipHnKklvB7H6Gocx+vSUonj6DoccRUZVoFjbP2NWKoMrrwxvu7GaBHZyIRV3Diq39iMw6HyGsecPQ4srNEsh3uaBC1KXqYPxQpKQ70y1/c0sUa8B5a2IgMRuDIMi/9+T5NbBedOke1cpcbWzQdDl+K50U3rHk58AderuTE5r1iLOc84oQ/dsdbZ08y5/vGO0E00tX8pSSOD9qhr9AIhTrRdvSBHdIABlGfpDf7GuLtEL24gki2kRTatTKoRO2jozYGDdxA46B00PZLNUAFUwzYpsigNu6/6KcJSwi6XfdI1z2Y+iMhWDVcb3TkgMwdkhiDNAZk5IDMS0RyQmQMyc0BmDsjMARkrBm8FwJK/rf6fF8KY2n89b6xbce+4RRL+E87vmP6NpkgyBDQ1GmNflgJhTwlMjEDjGYBdRNNGhB2TbyTmLF3kHJSbohDogqG7qTBuWYpaoqgi6kFQOUox+NakvK1uJF4p6JwG3r2lswyvJD3E02w6G4igBaOHY2LQ1NVLP3QZv+5b9Y8uCdUTT3vxOrHXrhMSyyLeIet8i4U7Y9DegG4jfJnITXM0I3RRVX29RE+YSP0fCXxHKPbfpAg4dZ8Dt1fQDUTZItRM7PI9sJiUB+rOxyJUwqZX6vcIMCWfwWrYvcqhJphJ+vut1BC6aFBd60qTSmnXHIvtZ8byn3HywNbrS/Q3zvWJsNsiyy5R89/q+75q1cN4o301A11cs12egYT0spXENaaUybuCahaMX6J//vPXTyTLIP2pav7COlDGnPsYLC6vU5Bd5x1Kuq7M41Fqv779put/iZKlR++1UXsWSBU7SJGd4aGcfGdDBpIWcw6Jmgqu0H8v/isG8gZLoEB92IfhTU3JdEn9rDXJSiWe/rKoIRFUSd5l8vxgTYNagS+Pu1Vbnb/vOg2bcEb/YKtYJk1JLdI1gr39l3CjBl1XSHo0uluDUxlY6RgmY1UY3D42Qvi0JFDOMtKh1By9SJTZPOFGlTaqUJJSXpFob6budRPD8hRLUYgcaNo7iu4zjg64mwGFuhMR5bXa6LZ9V5e1tgT6PW7Iobeas2SLRC/UX0N4wsJaPLuZp7CQy7oHRMOhhK7LytcweEHtAwSeT8ReUR5knwJOM0LdnIf63C8VgYY1XkvgzZDSSBKmL2TgygxcY5IZmgj5j/9Pt7OXQp6x/W7idRXG1NgSjOLu5dhS+SF4uPXXj09WpCUXm0PSrnl5RhIc7g0ehaPmgghds5G2RAqCcE8Rpkne0i8txjbHpuLYor4QOSRTjs3FwtgWF3HozTj2Ss8Hy+AVACxPrZcZRAdV8ukDMs/BRpocYlbDjhma8Qc9Jrn2uiCzGe9AF5IXcInWOBOg/PKCPlD2RN3jpqDVSuHtpJNCMxrlAR/fZBjT3zcO4Z7OxW6qYptHfv3+dV3maQDUhPqqNaamoNT5SmEbMn8pJ+6L6wT2kPfZKOZFkVdo/cXAjI2Xk+hOn2U/Vdc0daN8pGGFnBSOPsrfLUlnF/CUFSpWrStiv9Oo93G4maiQ3dxamW2ZkMvTcFSkXWxHLsLjGFeL5XHVf0641dKBWe213NV7LbdAU0I3i8Xi2C2WmOim2R21f+q2QWNibbjZ8F720XY9M4jlwVYEq9IO06eCE7qOJlS3DxvjbowJtQ+2h7e+Vr5iDhzdlX/cWyqGhHq1L4XLP4bjoVLjdyw2ttJVH04ltOrCN32NUsUJrfY6ntiC07v+nGXdcx/oIKq4At/sEkuK6yLL9jW3QWka28/6AMP3gh1E8qdNLQbNSLsHp4sO31Vo/1ejHYoRd+U0BkHJgVBlnUOKLraYp3qJEpD+5DtSEmcT5bChzq0U2b27cwQLs4Xl2FGvXqLfVVN/V239XTX2d8cKYmn4Ee3T5LQoyw6I8zwjIJBsY40dMq4/+/9pjAbgjySJFTCqqPmGCjznwMkO6KHzcaoNt/sKkaLhyIbICiGBH2eQ31AJnOIM3dw2/b4Sgp0bPJcvLGM0qiaGfvly7x4HDUtHC49h6HAxMobT5QpnmCZuiQbw+8xwin6u6DS9ysF0yjivG9aj0TiGdMNBiAltuSkpuNDXDJTn5uwTQ52w5PAPG4mD1Vv0kjg9KZwHCZwhStMcymMKvVmltwDal56BmtK1upqLfS0vHDEHSixhXWTxfJKaYjSnxCe0oe1My03hW0OEza3O6AKUZVEu4PdVC7qG6xm8pAPhNebfUY7SiU1r4/x6bVkfmKsuIaIX8Jh6u38+gDW41nc4tZ4NL8Wwul6XuhslG2Bfh5pr5QYAawuiCAm7WPNdec+BUbs1yqRnyWxHY5ZeS53yJhfNnYE139zbfdyAfFQPGvnmbu6dL+21kJuLkpsm0hsvJDzfT3v4zPfThuEZrqv8yLJiFysmVRKLYpBMidn0pfKvEpjTEJkvDK2ewPE3Xxg6VkDzhaHt8y4vDP0WeE3oGW7l/LvjLs4ulHPcWFoaeRWYfwcAAP//5CGUeg=="
}
