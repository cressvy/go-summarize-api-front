package main

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	openai "github.com/sashabaranov/go-openai"

	"github.com/unidoc/unipdf/v3/common/license"
	"github.com/unidoc/unipdf/v3/extractor"
	"github.com/unidoc/unipdf/v3/model"
)

func init() {
	err := license.SetMeteredKey("fuck you")
	if err != nil {
		panic(err)
	}
}

func main() {
	// set with default port 8080
	r := gin.Default()

	// load HTML file
	// to view the HTML file, go to http://localhost:8080
	r.LoadHTMLGlob("front/*")

	// GET /
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	// POST /summarize
	r.POST("/summarize", func(c *gin.Context) {
		// get file from form-data key "file"
		file, err := c.FormFile("file")

		// log file name
		fmt.Println(file.Filename)

		if err != nil {
			c.String(http.StatusBadRequest, "Error: %v", err)
			return
		}

		openFile, err := file.Open()

		if err != nil {
			c.String(http.StatusInternalServerError, "Error: %v", err)
			return
		}
		defer openFile.Close()

		pdfReader, err := model.NewPdfReader(openFile)
		numPages, err := pdfReader.GetNumPages()

		content := "summarize this:\n"
		for i := 0; i < numPages; i++ {
			pageNum := i + 1
			page, err := pdfReader.GetPage(pageNum)
			if err != nil {
				c.String(http.StatusInternalServerError, "Error: %v", err)
				return
			}

			ex, err := extractor.New(page)
			if err != nil {
				c.String(http.StatusInternalServerError, "Error: %v", err)
				return
			}

			pageText, err := ex.ExtractText()
			if err != nil {
				c.String(http.StatusInternalServerError, "Error: %v", err)
				return
			}

			content += pageText + "\n"

		}

		client := openai.NewClient("sk-proj-TecKHczjbuyM8ykQp8mvT3BlbkFJeeJpIOirRlJzb2YewbeD")
		resp, err := client.CreateChatCompletion(
			c,
			openai.ChatCompletionRequest{
				Model: openai.GPT3Dot5Turbo,
				Messages: []openai.ChatCompletionMessage{
					{
						Role:    openai.ChatMessageRoleUser,
						Content: content,
					},
				},
			},
		)
		if err != nil {
			c.String(http.StatusInternalServerError, "ChatCompletion error: %v", err)
			return
		}

		c.String(http.StatusOK, resp.Choices[0].Message.Content)

	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func readPDFWithUniPDF(file multipart.File) error {
	defer file.Close()

	// Open a new PDF reader
	pdfReader, err := model.NewPdfReader(file)
	if err != nil {
		return err
	}

	// Access and process PDF information using UniPDF functions
	// (Replace this with your specific logic for summarizing the PDF)
	numPages, err := pdfReader.GetNumPages()
	if err != nil {
		return err
	}
	fmt.Println("The uploaded PDF has", numPages, "pages.")

	// ... further processing or summarization using UniPDF functionalities ...

	return nil
}

func extractAndSummarizePdf(fileBytes *os.File) (string, error) {
	pdfReader, err := model.NewPdfReader(fileBytes)
	if err != nil {
		return "", err
	}

	numPages, err := pdfReader.GetNumPages()
	if err != nil {
		return "", err
	}

	var text string
	for i := 0; i < numPages; i++ {
		pageNum := i + 1
		page, err := pdfReader.GetPage(pageNum)
		if err != nil {
			return "", err
		}

		ex, err := extractor.New(page)
		if err != nil {
			return "", err
		}

		pageText, err := ex.ExtractText()
		if err != nil {
			return "", err
		}

		text += pageText + "\n"
	}

	text = "summarize this:\n" + text
	return text, nil
}
