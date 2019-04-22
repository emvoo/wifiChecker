# WiFi Checker

This simple application is written in Go for linux operating system (for now). The aim of the application is for parents trying to restrict child's usage of computer (wireless connectivity in particular) out of allowed hours. Application will run every 60 seconds (this is default and can be overridden) to check if computer is connected to the internet and will keep wifi disabled 7pm - 8am (those values can be overridden at application runtime)

## Usage

To run with default values:

        ./checker

Command to run the application and make a check every 30s, and disable wifi 3pm - 10am
        
        ./checker --interval=30 --from=15:00 --to=10:00

It's best to launch application at computer boot.

        NAME:
           checker - A new cli application
        
        USAGE:
           checker [global options] command [command options] [arguments...]
        
        VERSION:
           1.0.0
        
        DESCRIPTION:
           Application to run the scripts on 60 seconds (default, can be overridden) intervals.
        
        COMMANDS:
             help, h  Shows a list of commands or help for one command
        
        GLOBAL OPTIONS:
           --interval value, -i value  use to set how often the app should check wifi connection, default every 60s (default: 60)
           --from value, -f value      what time wifi should be available at (default: "08:00")
           --to value, -t value        what time should the wifi be disabled (default: "19:00")
           --help, -h                  show help
           --version, -v               print the version


Feel free to suggest or better yet make changes to the application to add functionality. Application is still in development and has not been unit tested yet.