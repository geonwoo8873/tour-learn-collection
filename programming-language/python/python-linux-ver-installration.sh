#!/bin/env bash

#### User is prompted for the Python version to install
PS3="Please select the Python version to install: "
options=("Python 3.10" "Python 3.11" "Python 3.12" "Other version" "Quit")
select opt in "${options[@]}"
do
    case $opt in
        "Python 3.10")
            echo "You chose to install Python 3.10"
            break
            ;;
        "Python 3.11")
            echo "You chose to install Python 3.11"
            break
            ;;
        "Python 3.12")
            echo "You chose to install Python 3.12"
            break
            ;;
        "Quit")
            break
            ;;
        "Other version")
            echo "You chose to install another version of Python"
            break
            ;;
        # `$REPLY` is the user's other inputed selections to replay output menu
        *) echo "Invalid option $REPLY";;
    esac
done