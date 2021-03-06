package main

import (
	"testing"

	"github.com/matryer/is"
)

func TestParse(t *testing.T) {
	is := is.New(t)
	patterns := []string{"./testdata/services/pleasantries"}
	parser := newParser(patterns...)
	parser.ExcludeInterfaces = []string{"Ignorer"}
	def, err := parser.parse()
	is.NoErr(err)

	is.Equal(def.PackageName, "pleasantries")
	is.Equal(len(def.Services), 2) // should be 2 services
	is.Equal(def.Services[0].Name, "GreeterService")
	is.Equal(def.Services[0].Comment, `GreeterService is a polite API.
You will love it.`)
	is.Equal(len(def.Services[0].Methods), 2)
	is.Equal(def.Services[0].Methods[0].Name, "GetGreetings")
	is.Equal(def.Services[0].Methods[0].NameLowerCamel, "getGreetings")
	is.Equal(def.Services[0].Methods[0].Comment, "GetGreetings gets a range of saved Greetings.")
	is.Equal(def.Services[0].Methods[0].InputObject.TypeName, "GetGreetingsRequest")
	is.Equal(def.Services[0].Methods[0].InputObject.Multiple, false)
	is.Equal(def.Services[0].Methods[0].InputObject.Package, "")
	is.Equal(def.Services[0].Methods[0].OutputObject.TypeName, "GetGreetingsResponse")
	is.Equal(def.Services[0].Methods[0].OutputObject.Multiple, false)
	is.Equal(def.Services[0].Methods[0].OutputObject.Package, "")

	is.Equal(def.Services[0].Methods[1].Name, "Greet")
	is.Equal(def.Services[0].Methods[1].NameLowerCamel, "greet")
	is.Equal(def.Services[0].Methods[1].Comment, "Greet creates a Greeting for one or more people.")
	is.Equal(def.Services[0].Methods[1].InputObject.TypeName, "GreetRequest")
	is.Equal(def.Services[0].Methods[1].InputObject.Multiple, false)
	is.Equal(def.Services[0].Methods[1].InputObject.Package, "")
	is.Equal(def.Services[0].Methods[1].OutputObject.TypeName, "GreetResponse")
	is.Equal(def.Services[0].Methods[1].OutputObject.Multiple, false)
	is.Equal(def.Services[0].Methods[1].OutputObject.Package, "")

	greetInputObject, err := def.Object(def.Services[0].Methods[0].InputObject.TypeName)
	is.NoErr(err)
	is.Equal(greetInputObject.Name, "GetGreetingsRequest")
	is.Equal(greetInputObject.Comment, "GetGreetingsRequest is the request object for GreeterService.GetGreetings.")
	is.Equal(len(greetInputObject.Fields), 1)
	is.Equal(greetInputObject.Fields[0].Name, "Page")
	is.Equal(greetInputObject.Fields[0].NameLowerCamel, "page")
	is.Equal(greetInputObject.Fields[0].Comment, "Page describes which page of data to get.")
	is.Equal(greetInputObject.Fields[0].OmitEmpty, false)
	is.Equal(greetInputObject.Fields[0].Type.TypeName, "services.Page")
	is.Equal(greetInputObject.Fields[0].Type.ObjectName, "Page")
	is.Equal(greetInputObject.Fields[0].Type.ObjectNameLowerCamel, "page")
	is.Equal(greetInputObject.Fields[0].Type.JSType, "object")
	is.Equal(greetInputObject.Fields[0].Type.TypeID, "github.com/pacedotdev/oto/testdata/services.Page")
	is.Equal(greetInputObject.Fields[0].Type.IsObject, true)
	is.Equal(greetInputObject.Fields[0].Type.Multiple, false)
	is.Equal(greetInputObject.Fields[0].Type.Package, "github.com/pacedotdev/oto/testdata/services")
	is.Equal(greetInputObject.Fields[0].Tag, `tagtest:"value,option1,option2"`)
	is.True(greetInputObject.Fields[0].ParsedTags != nil)
	is.Equal(greetInputObject.Fields[0].ParsedTags["tagtest"].Value, "value")
	is.Equal(len(greetInputObject.Fields[0].ParsedTags["tagtest"].Options), 2)
	is.Equal(greetInputObject.Fields[0].ParsedTags["tagtest"].Options[0], "option1")
	is.Equal(greetInputObject.Fields[0].ParsedTags["tagtest"].Options[1], "option2")

	greetOutputObject, err := def.Object(def.Services[0].Methods[0].OutputObject.TypeName)
	is.NoErr(err)
	is.Equal(greetOutputObject.Name, "GetGreetingsResponse")
	is.Equal(len(greetOutputObject.Fields), 2)
	is.Equal(greetOutputObject.Fields[0].Name, "Greetings")
	is.Equal(greetOutputObject.Fields[0].NameLowerCamel, "greetings")
	is.Equal(greetOutputObject.Fields[0].Type.TypeID, "github.com/pacedotdev/oto/testdata/services/pleasantries.Greeting")
	is.Equal(greetOutputObject.Fields[0].OmitEmpty, false)
	is.Equal(greetOutputObject.Fields[0].Type.TypeName, "Greeting")
	is.Equal(greetOutputObject.Fields[0].Type.Multiple, true)
	is.Equal(greetOutputObject.Fields[0].Type.Package, "")
	is.Equal(greetOutputObject.Fields[1].Name, "Error")
	is.Equal(greetOutputObject.Fields[1].NameLowerCamel, "error")
	is.Equal(greetOutputObject.Fields[1].OmitEmpty, true)
	is.Equal(greetOutputObject.Fields[1].Type.TypeName, "string")
	is.Equal(greetOutputObject.Fields[1].Type.Multiple, false)
	is.Equal(greetOutputObject.Fields[1].Type.Package, "")

	is.Equal(def.Services[1].Name, "Welcomer")
	is.Equal(len(def.Services[1].Methods), 1)

	is.Equal(def.Services[1].Methods[0].InputObject.TypeName, "WelcomeRequest")
	is.Equal(def.Services[1].Methods[0].InputObject.Multiple, false)
	is.Equal(def.Services[1].Methods[0].InputObject.Package, "")
	is.Equal(def.Services[1].Methods[0].OutputObject.TypeName, "WelcomeResponse")
	is.Equal(def.Services[1].Methods[0].OutputObject.Multiple, false)
	is.Equal(def.Services[1].Methods[0].OutputObject.Package, "")

	welcomeInputObject, err := def.Object(def.Services[1].Methods[0].InputObject.TypeName)
	is.NoErr(err)
	is.Equal(welcomeInputObject.Name, "WelcomeRequest")
	is.Equal(len(welcomeInputObject.Fields), 4)

	is.Equal(welcomeInputObject.Fields[0].Name, "To")
	is.Equal(welcomeInputObject.Fields[0].NameLowerCamel, "to")
	is.Equal(welcomeInputObject.Fields[0].OmitEmpty, false)
	is.Equal(welcomeInputObject.Fields[0].Type.TypeName, "string")
	is.Equal(welcomeInputObject.Fields[0].Type.Multiple, false)
	is.Equal(welcomeInputObject.Fields[0].Type.Package, "")
	is.Equal(welcomeInputObject.Fields[0].Example, "your@email.com")

	is.Equal(welcomeInputObject.Fields[1].Name, "Name")
	is.Equal(welcomeInputObject.Fields[1].NameLowerCamel, "name")
	is.Equal(welcomeInputObject.Fields[1].OmitEmpty, false)
	is.Equal(welcomeInputObject.Fields[1].Type.TypeName, "string")
	is.Equal(welcomeInputObject.Fields[1].Type.JSType, "string")
	is.Equal(welcomeInputObject.Fields[1].Type.Multiple, false)
	is.Equal(welcomeInputObject.Fields[1].Type.Package, "")
	is.Equal(welcomeInputObject.Fields[1].Example, "John Smith")

	is.Equal(welcomeInputObject.Fields[2].Example, float64(3))
	is.Equal(welcomeInputObject.Fields[2].Type.JSType, "number")

	is.Equal(welcomeInputObject.Fields[3].Example, true)
	is.Equal(welcomeInputObject.Fields[3].Type.JSType, "boolean")

	welcomeOutputObject, err := def.Object(def.Services[1].Methods[0].OutputObject.TypeName)
	is.NoErr(err)
	is.Equal(welcomeOutputObject.Name, "WelcomeResponse")
	is.Equal(len(welcomeOutputObject.Fields), 2)
	is.Equal(welcomeOutputObject.Fields[0].Name, "Message")
	is.Equal(welcomeOutputObject.Fields[0].NameLowerCamel, "message")
	is.Equal(welcomeOutputObject.Fields[0].Type.IsObject, false)
	is.Equal(welcomeOutputObject.Fields[0].OmitEmpty, false)
	is.Equal(welcomeOutputObject.Fields[0].Type.TypeName, "string")
	is.Equal(welcomeOutputObject.Fields[0].Type.Multiple, false)
	is.Equal(welcomeOutputObject.Fields[0].Type.Package, "")
	is.Equal(welcomeOutputObject.Fields[1].Name, "Error")
	is.Equal(welcomeOutputObject.Fields[1].NameLowerCamel, "error")
	is.Equal(welcomeOutputObject.Fields[1].OmitEmpty, true)
	is.Equal(welcomeOutputObject.Fields[1].Type.TypeName, "string")
	is.Equal(welcomeOutputObject.Fields[1].Type.Multiple, false)
	is.Equal(welcomeOutputObject.Fields[1].Type.Package, "")

	is.Equal(len(def.Objects), 8)
	for i := range def.Objects {
		switch def.Objects[i].Name {
		case "Greeting":
			is.Equal(len(def.Objects[i].Fields), 1)
			is.Equal(def.Objects[i].Imported, false)
		case "Page":
			is.Equal(def.Objects[i].TypeID, "github.com/pacedotdev/oto/testdata/services.Page")
			is.Equal(len(def.Objects[i].Fields), 3)
			is.Equal(def.Objects[i].Imported, true)
		}
	}

	// b, err := json.MarshalIndent(def, "", "  ")
	// is.NoErr(err)
	// log.Println(string(b))
}

func TestExtractExample(t *testing.T) {
	is := is.New(t)

	example, comment, err := extractExample(`
		This is a comment
		example: "With an example"
	`)
	is.NoErr(err)
	is.Equal(comment, "This is a comment")
	is.Equal(example, "With an example")

	example, comment, err = extractExample(`
		This is a comment
		example: true
	`)
	is.NoErr(err)
	is.Equal(comment, "This is a comment")
	is.Equal(example, true)

	example, comment, err = extractExample(`
		This is a comment
		example: 123
	`)
	is.NoErr(err)
	is.Equal(comment, "This is a comment")
	is.Equal(example, float64(123))

}
