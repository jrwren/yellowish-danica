package danica

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/robertkrimen/otto/parser"
)

func TestHelloWorld(t *testing.T) {
	/*bundlefile := `module.exports = {
		src: "src/main.js",
		dest: "dist/out.js"
	  };`*/
	src_main_js := `console.log("Hello world");`
	os.Mkdir("src", os.ModePerm)
	ioutil.WriteFile("src/main.js", []byte(src_main_js), os.ModePerm)

	err := BundleFile("src/main.js", "dist/out.js")
	assert.Nil(t, err)

	r, err := ioutil.ReadFile("dist/out.js")
	assert.Nil(t, err)
	assert.Equal(t, string(r), crazy_miguel_stuff+"\n"+src_main_js)
}

func TestFromContent(t *testing.T) {
	content := `
    var hello = require("./src/hello");
    var world = require("./src/world");
    console.log(hello() + " + " + world());
  `

	os.Mkdir("src", os.ModePerm)
	hello := `module.exports = () => "hello";`
	ioutil.WriteFile("src/hello", []byte(hello), os.ModePerm)
	world := `module.exports = () => "world";`
	ioutil.WriteFile("src/world", []byte(world), os.ModePerm)

	err := BundleContent(content, "dist/out.js")
	if err != nil {
		el, ok := err.(parser.ErrorList)
		if ok {
			for i := range el {
				t.Logf("%s", el[i])
			}
		}
	}
	assert.Nil(t, err)
	r, err := ioutil.ReadFile("dist/out.js")
	assert.Nil(t, err)
	assert.Equal(t, string(r), crazy_miguel_stuff+"\n"+content+"\n"+
		hello+"\n"+
		world)
}
