package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"mime/multipart"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

func uploadFileToDrive(ctx context.Context, file multipart.File, header *multipart.FileHeader) (*drive.File, error) {
	b, err := os.ReadFile("credentials.json")
	if err != nil {
		return nil, err
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, drive.DriveFileScope)
	if err != nil {
		return nil, err
	}
	client := getClient(config)

	srv, err := drive.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		return nil, err
	}

	// Create a new file metadata
	driveFile := &drive.File{
		Name: header.Filename,
	}

	// Upload the file
	driveFile, err = srv.Files.Create(driveFile).Media(file).Do()
	if err != nil {
		return nil, err
	}

	if _, err := srv.Permissions.Create(driveFile.Id, &drive.Permission{
		Role: "reader",
		Type: "anyone",
	}).Do(); err != nil {
		return nil, err
	}

	getDriveFile, err := srv.Files.Get(driveFile.Id).Fields("name, size, thumbnailLink, mimeType, webViewLink").Do()
	if err != nil {
		return nil, err
	}

	return getDriveFile, nil
}

// Retrieve a token, saves the token, then returns the generated client.
func getClient(config *oauth2.Config) *http.Client {
	// The file token.json stores the user's access and refresh tokens, and is
	// created automatically when the authorization flow completes for the first
	// time.
	tokFile := "token.json"
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(tokFile, tok)
	}
	return config.Client(context.Background(), tok)
}

// Request a token from the web, then returns the retrieved token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	log.Println("Input here")
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code %v", err)
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web %v", err)
	}
	return tok
}

// Retrieves a token from a local file.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}