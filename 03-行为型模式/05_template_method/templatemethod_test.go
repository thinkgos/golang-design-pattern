package templatemethod

func ExampleTemplateDownloader() {
	var downloader = newTemplate()

	downloader.Download("tpl://example.com/abc.zip")
	// Output:
	// prepare downloading
	// default download
	// default save
	// finish downloading
}

func ExampleHTTPDownloader() {
	var downloader = NewHTTPDownloader()

	downloader.Download("http://example.com/abc.zip")
	// Output:
	// prepare downloading
	// download http://example.com/abc.zip via http
	// http save
	// finish downloading
}

func ExampleFTPDownloader() {
	var downloader = NewFTPDownloader()

	downloader.Download("ftp://example.com/abc.zip")
	// Output:
	// prepare downloading
	// download ftp://example.com/abc.zip via ftp
	// default save
	// finish downloading
}
