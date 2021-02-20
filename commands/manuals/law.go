package manuals

var lawManual Manual = Manual{
	Summary: "law file",
	Content: `The "Law File" is a YAML file that describes the possible values 
of a file header. Here is an example Law File:

expected: ^// © ABC Company 1996. All rights reserved.$

Explanation:

	expected  Regular expression of the expected file header.

The Law File should be located in the current directory from which lawyer is
executed with the name "law" and the "yml" or "yaml" extension. 

Also note that YAML has multi-line support for file headers that require 
multiple lines or even newlines after multiple lines. See the YAML specfication 
for more details.`,
}
