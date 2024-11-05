package main

import (
	"fmt"
	"strconv"
	"syscall"
	"time"
	"toggl_to_xlsx/integration/excel"
	toggl "toggl_to_xlsx/integration/toggl"
	togglResponse "toggl_to_xlsx/integration/toggl/responses"

	"golang.org/x/term"
)

type Report struct {
	ProjectName string    `json:"project_name"`
	Start       time.Time `json:"start"`
	Stop        time.Time `json:"stop"`
	Duration    int       `json:"duration"`
	Description string    `json:"description"`
	At          time.Time `json:"at"`
}

var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"
var Magenta = "\033[35m"
var Cyan = "\033[36m"
var Gray = "\033[37m"
var White = "\033[97m"
var Transparent = "\033[8m"

func main() {
	fmt.Print(Green + `
___________________________________________________________________________________________________
   _____                  _  __  __  __  ____  __    ___                                         
  /__   \___   __ _  __ _| | \ \/ / / / / _\ \/ /   / __\___  _ ____   _____ _ __ ___  ___  _ __ 
    / /\/ _ \ / _' |/ _' | |  \  / / /  \ \ \  /   / /  / _ \| '_ \ \ / / _ \ '__/ __|/ _ \| '__|
   / / | (_) | (_| | (_| | |  /  \/ /____\ \/  \  / /__| (_) | | | \ V /  __/ |  \__ \ (_) | |   
   \/   \___/ \__, |\__, |_| /_/\_\____/\__/_/\_\ \____/\___/|_| |_|\_/ \___|_|  |___/\___/|_|   
              |___/ |___/                                                                        
___________________________________________________________________________________________________
` + "\n\n" + Reset)

	fmt.Print("Type your email: ")
	email := ""
	fmt.Scanln(&email)

	fmt.Print(Transparent + "Type your password: ")
	password, error := term.ReadPassword(int(syscall.Stdin))

	if error != nil {
		fmt.Print(Red + "Password is invalid" + Reset)
	}

	togglClient := toggl.Toggl{Email: email, Password: string(password)}
	workspaces, error := toggl.GetWorkspaces(&togglClient)
	if error != nil {
		fmt.Println("\n" + Red + error.Error() + Reset)
		syscall.Exit(1)
	}
	fmt.Println()
	for i, workspace := range workspaces {
		fmt.Println(Cyan + "[ " + strconv.Itoa(i) + " ] - " + workspace.Name + Reset)
	}

	fmt.Print("Select your workspace: ")
	var selectedWorkspaceIndex int
	fmt.Scanln(&selectedWorkspaceIndex)
	workspaceId := workspaces[selectedWorkspaceIndex].Id

	projects, error := toggl.GetProjects(workspaceId, &togglClient)
	if error != nil {
		fmt.Println("\n" + Red + error.Error() + Reset)
		syscall.Exit(1)
	}

	fmt.Println()
	for i, project := range projects {
		fmt.Println(Cyan+"[ "+strconv.Itoa(i)+" ] - ", project.Name+Reset)
	}

	fmt.Println("Select your project: ")
	var selectedProjectIndex int
	fmt.Scanln(&selectedProjectIndex)
	project := projects[selectedProjectIndex]

	var startDate = ""
	var endDate = ""
	fmt.Print("Type report start date (YYYY-MM-DD): ")
	fmt.Scanln(&startDate)
	fmt.Print("Type report end date (YYYY-MM-DD): ")
	fmt.Scanln(&endDate)

	timeEntries, error := toggl.GetTimeEntries(startDate, endDate, &togglClient)
	var filteredTimeEntries []togglResponse.TimeEntriesResponse

	for _, timeEntry := range timeEntries {
		if timeEntry.ProjectId == project.Id {
			filteredTimeEntries = append(filteredTimeEntries, timeEntry)
		}
	}

	if error != nil {
		fmt.Println("\n" + Red + error.Error() + Reset)
		syscall.Exit(1)
	}

	fmt.Println()
	var report []map[string]any
	for _, timeEntry := range filteredTimeEntries {
		report = append(report, map[string]any{"Date": formatDate(timeEntry.Start), "Duration": formatDuration(timeEntry.Duration), "Project": project.Name, "Description": timeEntry.Description})
	}
	filename := ""
	fmt.Print("Type your filename (Without extension): ")
	fmt.Scan(&filename)
	excel.MakeSheet(filename, report)

	fmt.Println("\n" + Green + "Done!" + Reset)
}

func formatDate(date time.Time) string {
	year, month, day := date.Date()

	return fmt.Sprintf("%02d/%02d/%02d", day, month, year)
}

func formatDuration(duration time.Duration) string {
	duration = duration.Round(time.Second)
	hour := duration / time.Hour
	duration -= hour * time.Hour
	minute := duration / time.Minute
	duration -= minute * time.Minute
	second := duration / time.Second
	return fmt.Sprintf("%02d:%02d:%02d", hour, minute, second)
}
