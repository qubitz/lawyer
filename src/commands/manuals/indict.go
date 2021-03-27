package manuals

var indictManual Manual = Manual{
	Description: "verify files have the expected header",
	Content: `Reads the top 10 lines (the header) of every file found matching
<file_pattern> and compares them to the regular expression of
"expected" inside the law file. Enter "lawyer help law" for more
information on the law file.

Usage:

	lawyer indict <file_pattern>

Arguments:

	<file_pattern>  file(s) to indict
	
If the header does not match the expected header, the header
is printed as evidence and the file is marked guilty.`,
}
