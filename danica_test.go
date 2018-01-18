package danica

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
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
    const hello = require("./src/hello");
    const world = require("./src/world");
    console.log(hello() + " + " + world());
  `

	os.Mkdir("src", os.ModePerm)
	ioutil.WriteFile("src/hello", []byte(`module.exports = () => "hello";`), os.ModePerm)
	ioutil.WriteFile("src/world", []byte(`module.exports = () => "world";`), os.ModePerm)

	err := BundleContent(content, "dist/out.js")
	assert.Nil(t, err)
	//TODO: this next
}