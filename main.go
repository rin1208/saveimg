package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"github.com/urfave/cli"
	"github.com/PuerkitoBio/goquery"
)

func main() {
	app := cli.NewApp()
    app.Name = "ImageSaveTool"
    app.Usage = "savetool"
    app.Version = "1.0.0"
    app.Action = func(c *cli.Context) error {
	if c.Bool("image") {
		saveimg()
	}  else if c.Bool("site"){
		savesite()
	} else {
		var a string
		fmt.Print("site URL or image URL? :")
		fmt.Scan(&a)
		if a == "site" {
			savesite()
		} else if a == "image" {
			saveimg()
		} else if a != "site" && a != "image" {
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

func savesite() {
	var siteurl string
	defer fmt.Println("Save completed!!")
	var result []*url.URL
	fmt.Print("Please enter URL :")
	fmt.Scan(&siteurl)
	confirmation := confirmurl(siteurl)
	if confirmation == true {
		doc, _ := goquery.NewDocument(siteurl)
		doc.Find("img").Each(func(_ int, s *goquery.Selection) {
			target, _ := s.Attr("src")
			base, _ := url.Parse(siteurl)
			targets, _ := url.Parse(target)
			result = append(result, base.ResolveReference(targets))
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
}

func saveimg() {
	var siteurl string
	defer fmt.Println("Save completed!!")
	fmt.Print("Please enter URL : ")
	fmt.Scan(&siteurl)
	confirmation := confirmurl(siteurl)
	if confirmation == true {
		_, err := exec.Command("wget", siteurl).Output()
		if err != nil {
			log.Fatal("Save failed...", err)
		}
	}
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

