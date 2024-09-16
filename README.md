# Explanation of the Code:
- Concurrency with Goroutines: The main.go launches a Goroutine for each file using the processFile function.
- WaitGroup: The sync.WaitGroup ensures the main program waits until all files are processed before closing the channels.
- Channels: Two channels are used:
    - resultChan: To send processed data (in this case, the name field from each JSON).
    - errChan: To send any errors encountered during the processing.
- Error Handling: Errors are handled and logged via the errChan, and processing continues for other files.
- File Reading and Processing: The file_reader.go is responsible for reading files, and processor.go processes the JSON and sends the result or error.


# How to run the project
- Navigate to the project directory.
- run => go run .
