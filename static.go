// Code generated by "esc -o static.go -pkg main static"; DO NOT EDIT.

package main

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"sync"
	"time"
)

type _escLocalFS struct{}

var _escLocal _escLocalFS

type _escStaticFS struct{}

var _escStatic _escStaticFS

type _escDirectory struct {
	fs   http.FileSystem
	name string
}

type _escFile struct {
	compressed string
	size       int64
	modtime    int64
	local      string
	isDir      bool

	once sync.Once
	data []byte
	name string
}

func (_escLocalFS) Open(name string) (http.File, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	return os.Open(f.local)
}

func (_escStaticFS) prepare(name string) (*_escFile, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	var err error
	f.once.Do(func() {
		f.name = path.Base(name)
		if f.size == 0 {
			return
		}
		var gr *gzip.Reader
		b64 := base64.NewDecoder(base64.StdEncoding, bytes.NewBufferString(f.compressed))
		gr, err = gzip.NewReader(b64)
		if err != nil {
			return
		}
		f.data, err = ioutil.ReadAll(gr)
	})
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (fs _escStaticFS) Open(name string) (http.File, error) {
	f, err := fs.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.File()
}

func (dir _escDirectory) Open(name string) (http.File, error) {
	return dir.fs.Open(dir.name + name)
}

func (f *_escFile) File() (http.File, error) {
	type httpFile struct {
		*bytes.Reader
		*_escFile
	}
	return &httpFile{
		Reader:   bytes.NewReader(f.data),
		_escFile: f,
	}, nil
}

func (f *_escFile) Close() error {
	return nil
}

func (f *_escFile) Readdir(count int) ([]os.FileInfo, error) {
	return nil, nil
}

func (f *_escFile) Stat() (os.FileInfo, error) {
	return f, nil
}

func (f *_escFile) Name() string {
	return f.name
}

func (f *_escFile) Size() int64 {
	return f.size
}

func (f *_escFile) Mode() os.FileMode {
	return 0
}

func (f *_escFile) ModTime() time.Time {
	return time.Unix(f.modtime, 0)
}

func (f *_escFile) IsDir() bool {
	return f.isDir
}

func (f *_escFile) Sys() interface{} {
	return f
}

// FS returns a http.Filesystem for the embedded assets. If useLocal is true,
// the filesystem's contents are instead used.
func FS(useLocal bool) http.FileSystem {
	if useLocal {
		return _escLocal
	}
	return _escStatic
}

// Dir returns a http.Filesystem for the embedded assets on a given prefix dir.
// If useLocal is true, the filesystem's contents are instead used.
func Dir(useLocal bool, name string) http.FileSystem {
	if useLocal {
		return _escDirectory{fs: _escLocal, name: name}
	}
	return _escDirectory{fs: _escStatic, name: name}
}

// FSByte returns the named file from the embedded assets. If useLocal is
// true, the filesystem's contents are instead used.
func FSByte(useLocal bool, name string) ([]byte, error) {
	if useLocal {
		f, err := _escLocal.Open(name)
		if err != nil {
			return nil, err
		}
		b, err := ioutil.ReadAll(f)
		_ = f.Close()
		return b, err
	}
	f, err := _escStatic.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.data, nil
}

// FSMustByte is the same as FSByte, but panics if name is not present.
func FSMustByte(useLocal bool, name string) []byte {
	b, err := FSByte(useLocal, name)
	if err != nil {
		panic(err)
	}
	return b
}

// FSString is the string version of FSByte.
func FSString(useLocal bool, name string) (string, error) {
	b, err := FSByte(useLocal, name)
	return string(b), err
}

// FSMustString is the string version of FSMustByte.
func FSMustString(useLocal bool, name string) string {
	return string(FSMustByte(useLocal, name))
}

var _escData = map[string]*_escFile{

	"/static/.DS_Store": {
		local:   "static/.DS_Store",
		size:    6148,
		modtime: 1537079402,
		compressed: `
H4sIAAAAAAAC/+yYMarDMBBEZ/RdGH6jMqWaHCA3EMY5QS6QIldIr6MHswMRiZLahHlgnsCWDTbs7hgA
l/vtBGQAM8I4Ysis443Ural7GGOMMWbfUP3936/CGDOoD0WucgtT55M8veYAxtbNVW5h6rokT/IsZ7nI
VW5hFS0qfFBPphIKlUJY5OrvaMyIv1De+v/5c/43xvxyn5/Wy7o8A8FwELh2629DQIqfhYdurwcBY3bG
IwAA//9qAIhtBBgAAA==
`,
	},

	"/static/index.html": {
		local:   "static/index.html",
		size:    9537,
		modtime: 1537069872,
		compressed: `
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
pfTv/wUAAP//M0VQVEElAAA=
`,
	},

	"/static/js/Api.js": {
		local:   "static/js/Api.js",
		size:    1448,
		modtime: 1537055387,
		compressed: `
H4sIAAAAAAAC/4RUUWvbMBB+96849ORsRk7aFYqzDMb2NMZWKKOwMYYqX5LbFMlIct2u+L8P2U7jOArT
i+3T992dvvtkqYRz8L4ieE4AAKTRzttaemPT2RALS6GH2ipYAWtckefKSKG2xvnien49zxvHli/YPJ+i
F3O+uOJX/E0E7bfkuDPyD3pYgcYG7vD+tvtOa6uykILNonhu9A6dExuEFaxrLT0ZDSk+oPbj5juafZpE
9sey6GoVan+6/fqFV8I67FPwUngxqrxfDenSNLxSwq+N3XHS4fEBtbdCfXNoXdrnnHBbkMLLLaQ4i7QS
lDcKuTKblJF+EIpKIC3Nbkd6A6EXBq8BIw2NqWc7b5PjtzbpHg51eYOmUpjuD5RB1QVCl8m4hoc95I5K
v4XVyzd39Bd5E6LLA2dtLKQ9kYD0KO20ew9VyNbt/6CfywggHAlWEeFEN/YCWPBkeGXZCaYST8qIsojQ
w6od2l9UFlBxKrMoZCd+G1vAYj4/s0867Hfu3KC/GYT5bvRB2E61DCr+ODvJ0R4P61iBYFPnLenNx16F
zqp9hNZPaWTcY08cqBPQ+C4FI8SRbZKcJUwaOcr+zPrRsAKYFIrurfDIMmDDNFgBz8xJ1MgKWLQHBWZD
8aHwf+R8/DzMffq/qoT1I4/2ns3h8nAyWkP6Qoe3HWNqT4u+thouR7QWUDmcsN/19V7BxbkMi5MMcdzF
8uSqtv8CAAD//xdeVcGoBQAA
`,
	},

	"/static/js/Histogram.js": {
		local:   "static/js/Histogram.js",
		size:    1037,
		modtime: 1537069670,
		compressed: `
H4sIAAAAAAAC/4xSXQ/SMBR93684YEIGzGUzUUzmfNYnjTwaHsZoWZO6Nm2FTbL/bgrdB2Ogfdlyzz0f
9+bmPNMaX5g24qiyX7h4AHBQ2dmXREhOlq5kHycGnGmDFJcm6cpUKPgWkxAUPa1rsI9RfybDE9Nsb9Fc
lIaVv0ky7oI/sxZhkelv5/K7EpIoU/syrJbDKF0kps1PGVY7pIiTO7gB4Zq84qzXI4bX/91Pd8oUpMsC
Vl4lxnFs9sno7mc6PzEoCDsWdqm3YK59l0w254ILNdpa6+6EPqfYTHnZd6UjxfwNpVEURfNHF7e3e8H3
WCzaoJ/+T/7j/pX8v/hRTumHeIrvPQ5vCqZDe7Q/SG58mSlNvpam3zzWiANsosB+3BQrvI2xwrvgZrp8
fgqDg+g8qgB1gDM7mKJVbIUGo1VYpTgqdtiyP6Q3qKfLV7VpqM08xDowN1VIGedbU3OCtD2RMfwst5u8
8Zq/AQAA//8UXpnTDQQAAA==
`,
	},

	"/static/js/Person.js": {
		local:   "static/js/Person.js",
		size:    2093,
		modtime: 1537065617,
		compressed: `
H4sIAAAAAAAC/5xVUU/bMBB+76+48YDSpQ30YS8rYZqQkCYxCW3TRB9N4jYnHLuyz7Qd8N8nJ2kaB4cK
+lLJ9913n893XzLBjIFbro2S8DQaAQBkShrSNiOlo+0EdmN4qs7djwo0CUqkO0hhO399voAUdr3zbQC7
C+AyVa4FJ/6LyVyVkAJpy3sYobIHnkMKSyZMP/iIBu8FH4hm1lBFGwreK6bzQWIs2crRSitEL1Qwc6vW
dj2QSZqh/E2MMBvidkVPrOF6egIxRHG8QZmrTYL5lbKSuB73MnI0pPHeEip5ozLm/iGF8zdg1yjRFIP3
23BJf9Q1akNX1UR4sJd6LoRi+Q/XiGjNqOhOheAEhoule7ICzYH8kWnAcuU6xzdQJ3dug+UqMdo1xjH6
50q6ek6IlVl1w6hb0v1cyfZpsFwdCF66ujdMPHzXyso86rZkAmvBaKl0Od5PflV7CR4MTk+74+Nh9/hm
yj+l6fAD9cX3ci/el3nYrTieB6NnZw1znMJPRkWiq62KxnAJ58kX+AYz+NqdmbZ3wIXhXW2XH9U2nX5U
23RI3OitJqbvfYBju+IbkD9dAUlN656ODkjtoM/Pey/0AouQ0tZID/nzMGjXBS18UFBg/SKNETDxUNtv
1G7HKHz3VzvWz+uUqT4pzgPu9uJWnGr8jcqa9kxgdtjJxOA/nmwwpwKmMOuYRsu1GObaBbgKjquCemRt
W522wCfKlfFM0KuVWa25pAmUKCdQsm3wzn+ZgP2cVy7kjfwYPkM0vKDT2RhiaArNPZdqqC/TurLmZLV8
DW23ucFfOLVH4E2wztjf/+V/AAAA//9zNvgNLQgAAA==
`,
	},

	"/static/js/Platform.js": {
		local:   "static/js/Platform.js",
		size:    3246,
		modtime: 1537068297,
		compressed: `
H4sIAAAAAAAC/6xW32/bNhB+919xeSmk1NaSdHuZoA6NhwEFOiBIt7VYEAy0dYqI0KRBUnHSwP/7QFI/
SFqyB6x+SCTdd3ffkXcfuWZEKbhhRFdCbuB1NgMAWAuutGzWWshE0W+Ywqv9bn66piozH6EA8y8PLVsU
W2Zsd/fOsnchFeovhD0mJVVa0lWjqeB+2EpISGxeoEC5H8yH9exgi1IJDoWPvKP3eQB1oGxH2OMHKRpe
Bvnn1jcdXPY+Y1KWNzZqYolOlOm9ZRVlGmVSNXxtwkOCLOYuUTeSQ3KGLFsJIkss/fxpPuQxtXH4AX4c
7K7yb4Ljh9b4bsx4bY3n457LwTOyroQsUZrIvxNdZxUTQib9dmc7WuraOKb5qKPJ2oU4h6tx0DIA+Rxs
oooy9rfgmNga53Ax79BpfgR53cO6h+uj+GUP6x6WaUxlQx7xL6roimGSBq1MuZmVJXItCftToVRJY/4G
XcJQQzsl1pgx5A+6hgVcxiujkFVtJ8UmyvXH0gwa6o9co3wizGuvIKH5PREJa9FwDQX8lIe2XU0Zgh1n
eA8X8OZNC30PFweBugqajv6d8bvPD1F2bg10C6KypUxMbfejFSTbjJZQFAU0mYn+Dy2n0HaIs3WjtNiY
NZIN5keAO+T6D/EblUovrbKZFBvKhbTp3l0ec/al4ZNYE7vIcQD4xRX5gPqWPtT6lvBSbL4mKfzcGz5h
NXw/lpFyqr9CMZF73HM/O/3FbNZicehvNzw27MNNNfvjGrcoXGfEUXaUl2KXKU20ae/29bN5/ayJ1L96
tYxwYEhk38y2wdNJQvt5F15iJVHVt0RjOI3xenuEbVsSqTuZDmUMFr5IGaxE1TDd6Z87L9yjDZ6kcG7j
pfDWn+JW0513zC1qkv9K7jQx63sOV/AWxjkG5U1zDHTO4/e9tcnUQRhrU43OstXMWuyWrYRdXYT2+I5w
Qm3ae4Kl2AMPrghdz59tsyfHbUqLesBRIQpqrAhTE7ih0EUBI7q0n43S9NZnaj7NbyWRPOYnpGJk8gf2
oyt6cnZP6MMXQjXlD/9n3vuDnM+BYaXnIM2IBf3mdt4NU+HssLBo7wRrD0Q+egC6CM9QWK9uxtyVKB43
mydaB+f/Et6lYsdh7mu0NeSHG3Kwus9wZjbenN8v/eMEKhaWBVwOfnF2Yw3ijHWAVa3u5s1xBzf2JXme
w8tIK/i3422j6sT5jiD54bEU38r3s38DAAD//7dO8RiuDAAA
`,
	},

	"/static/js/Size.js": {
		local:   "static/js/Size.js",
		size:    157,
		modtime: 1537018464,
		compressed: `
H4sIAAAAAAAC/0rOSSwuVgjOrEpVqOZSUFBQSM7PKy4pKk0uyS/SqNBRqNRRKM9MKcnQUchIzUzPKNGE
KgOBkozMYr0KBVuFCmtUsUoFW4VKNDGwKQq2ENPQ5CBGK9hC7YDI1nLVAgIAAP//AnTKup0AAAA=
`,
	},

	"/static/js/Train.js": {
		local:   "static/js/Train.js",
		size:    2532,
		modtime: 1537063785,
		compressed: `
H4sIAAAAAAAC/7RWTU/cMBC951cMN6ewAVpx2hqpPVRCaiVUekBCHNxkQqwaO7Kd/RDKf69sh2ziZLfb
quwp4xnPPL83Y28umDHwQzMu4SVJAAByJY3VTW6VJn7B/bi8P+sNY1U9MFVjB9Z297nmha12ZoX8qbLe
TOGlX7YVNxmX90BdleV43ZcCGkpGPlcXqC8feTazubZAYRuteYhAA9TIF/AC7YBH3hpVLRAoPDxOMDPL
89vY3wZ6mdZ8hcQxkIz3saL45JxMGJIud04vCBgUJVAfuYxcXNqbwpGE9kZa1CsmSNnI3HIlgQy59uqh
KLMNnFK4XJ6fj1y8BNK5KaUh0jMf5/C1BTLd1/MY0uUkas1lodaeE/Q8e/POmb7r7qyqayxG+9reas9e
d2gsNZrqO7OYjvgssGbaBj6PZ+wgW3vpGjlKpYGEdDVwGQJDV8zSFSJRGyWBDqMf6scpbyFwtnCb/FE0
NxNzIF7jwsdoRPbIOiPpAHrU/seK/hVL+++K/1RMF6O2/g8TMoKzYhqe1QrdlpIJg8tkKr0L4scLH2nO
5zTP/NF8WasbPCj7G2Fw3XRSZ0Llv7CYS+V+r9xMQe6OsoWF79xZt6viQlzDLq72lQmJApSD1VpAYbDP
+hEWl0fm/MZslWkmC/VMUriGi+xqT4nk8MrMUJ54mt7k6vzs+oTLpy9cclNNrs+/mabRqzMchH1vWR8g
0EKuGumeyPcXO/zrigt0V6NzXcPFZLzcxk1PvmpkQcY6vAMyeJ0X8CFN4TS+Bl2W7TFZund8kGaSp7+W
Ja7h1htkcwbbSJXuUrbh4XLMzLZlF3agcbuIQ+M+oT+rG1ORsDONzuC5Xix2Gdruz0WbJO3vAAAA//+G
wjqA5AkAAA==
`,
	},

	"/static/js/platform-drawer.js": {
		local:   "static/js/platform-drawer.js",
		size:    48,
		modtime: 1537066309,
		compressed: `
H4sIAAAAAAAC/0orzUsuyczPU0gpSiwPyEksScsvytUoKUrMzNNRSC/KTAnOrErVUUguqdBUqObiqgUE
AAD//83LGLQwAAAA
`,
	},

	"/static/js/train-drawer.js": {
		local:   "static/js/train-drawer.js",
		size:    1118,
		modtime: 1537066446,
		compressed: `
H4sIAAAAAAAC/6ST4W6DIBCA//MUl+5PbVmn1ixNjI+wP+tegCBUEqKJYyts2bsvAkOd07ZZmjTeHX5+
dwB/q6kSTQ1lS84vLRH1WnX/GE6tKI/ig2GgSmOgpH4nrxF8IgAAyRRoKMCu3ek8JE1Imj55FqWqQsFG
fbFi4lSpUHVhjmw9SD0zqtYxBgNbSLAHbiHFXdSbru72vPutrHSUz0OyGyB/UDQG4wnYdzBGlLSkLF3w
0K6TUUP3nYufh30eSWWPhwFxBvlEVLXjsmnatWM+wD6yssk/RAfUARY2kN7Cnpmjm0Ay2QXO4zhOlszC
Vl58eaavn7Enixo8I4cLA/oFyq4AIfSF0PT6XXO8aCOb1nH8fdSwKcICJ2mmKWc5SfsTN8rbAlV6x4WU
R2Ukg8J9Nx+V5nSjvGvvOwAA//+zOhZlXgQAAA==
`,
	},

	"/": {
		isDir: true,
		local: "",
	},

	"/static": {
		isDir: true,
		local: "static",
	},

	"/static/js": {
		isDir: true,
		local: "static/js",
	},
}
