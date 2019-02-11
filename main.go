package main

import (
    "github.com/urfave/cli"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"github.com/PuerkitoBio/goquery"
)

func main() {
    app := cli.NewApp()
    app.Name = "ImageSaveTool"
    app.Usage = "savetool"
    app.Version = "1.0.0"
    app.Action = func(c *cli.Context) error {
	if c.Bool("image") {
		defer fmt.Println("Save completed!!")
			var siteurl string
            fmt.Print("Please enter URL : ")
            fmt.Scan(&siteurl)
            confirmation:= confirmurl(siteurl)
            if confirmation == true {
                _, err := exec.Command("wget", siteurl).Output()
                if err != nil {
                    log.Fatal("Save failed...", err)
                }
            }
	}  else if c.Bool("site"){
		defer fmt.Println("Save completed!!")
        var result []*url.URL
   		var siteurl string
	    fmt.Print("Please enter URL :")
        fmt.Scan(&siteurl)
        confirmation:= confirmurl(siteurl)
        if confirmation  == true {
            doc, _ := goquery.NewDocument(siteurl)
            doc.Find("img").Each(func(_ int, s *goquery.Selection) {
                urll, _ := s.Attr("src")
                base, _ := url.Parse(siteurl)
                urls, _ := url.Parse(urll)
                result = append(result, base.ResolveReference(urls))
            })
            for _, b := range result {
                a := b.String()
                log.Println(a)
                _, err := exec.Command("wget", a).Output()
                if err != nil {
                    log.Fatal("Save failed...", err)
                    }
                }
            }
	} else {
	var siteurl string
	var alice string
	fmt.Print("site URL or image URL? :")
	fmt.Scan(&alice)
	if alice == "site" {
		defer fmt.Println("Save completed!!")
		var result []*url.URL
		fmt.Print("Please enter URL :")
		fmt.Scan(&siteurl)
		confirmation:= confirmurl(siteurl)
		if confirmation  == true {
			doc, _ := goquery.NewDocument(siteurl)
			doc.Find("img").Each(func(_ int, s *goquery.Selection) {
				urll, _ := s.Attr("src")
				base, _ := url.Parse(siteurl)
				urls, _ := url.Parse(urll)
				result = append(result, base.ResolveReference(urls))
			})
			for _, b := range result {
				a := b.String()
				log.Println(a)
				_, err := exec.Command("wget", a).Output()
				if err != nil {
					log.Fatal("Save failed...", err)
					}
				}
			}
		} else if alice  == "image" {
			defer fmt.Println("Save completed!!")
			fmt.Print("Please enter URL : ")
			fmt.Scan(&siteurl)
			confirmation:= confirmurl(siteurl)
			if confirmation == true {
				_, err := exec.Command("wget", siteurl).Output()
				if err != nil {
					log.Fatal("Save failed...", err)
				}
			}
		} else {
			fmt.Println("Please enter either")
		}
	}
	return nil
  }

	app.Flags = []cli.Flag{
    cli.BoolFlag {
      Name: "image, i",
      Usage: "Save from image URL",
    },
    cli.BoolFlag {
      Name: "site, s",
      Usage: "Save from website",
    },

}
    app.Run(os.Args)
}

func confirmurl(url string) bool {
    _, err := http.Get(url)
    if err != nil {
        fmt.Println("URL that does not exist")
        os.Exit(0)
    }
    use := true
    return use
}
