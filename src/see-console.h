#ifndef PUFFERSTARTER_CLI_SEE_CONSOLE_H
#define PUFFERSTARTER_CLI_SEE_CONSOLE_H

#include <cstdio>
#include <cstdlib>
#include <fstream>
#include <iostream>
#include "menu_and_authenticaton.h"

void printConsole(std::string ip, std::string token, int fromMenu, std::string serverID) {
    const char* tempLocation;
    std::string rawTempLocation;
    
    #ifdef _WIN32
        rawTempLocation = std::getenv("TEMP");
        rawTempLocation += "\\pufferstarter\\tempfile.tmp";
        tempLocation = rawTempLocation.c_str();
    #else
        rawTempLocation = "/tmp/pufferstartertempfile.tmp";
        tempLocation = rawTempLocation.c_str();
    #endif
    
    std::cout << "Console output running through 'more'. Use 'Q' to quit\n";

    std::string getStaticConsoleCommand = "curl -X GET -H \"Content-Type: application/json\" -H \"Authorization: Bearer ";
    std::string getServerConsoleCommand = getStaticConsoleCommand + token + "\" " + ip + "/daemon/server/" + serverID + "/console" + " -s";
    const char* command = getServerConsoleCommand.c_str();
    std::string rawOutput = executeCommand(command);
    sleep(3);
    
    std::ofstream tempfile(tempLocation);
    std::string output = rawOutput.erase(0, 28);
    int cutFrom = output.length() - 2;
    output = output.erase(cutFrom, 2);

    tempfile << output;
    tempfile.close();

    #ifdef _WIN32
        std::string rawMoreCommand = "more " + rawTempLocation;
        const char* moreCommand = rawMoreCommand.c_str();
        std::system(moreCommand);
    #else
        std::string rawLessCommand = "less " + rawTempLocation;
        const char* lessCommand = rawLessCommand.c_str();
        std::system(lessCommand);
    #endif
    std::remove(tempLocation);
    
    
    /* std::string rawOutput = executeCommand(command);
    std::string output = rawOutput.erase(0, 28);
    int cutFrom = output.length() - 2;
    output = output.erase(cutFrom, 2);
    std::cout << output; */

    if (fromMenu == 0) {
        exit(0);
    }
    menuReturn();
}


void getConsoleInfo(std::string ip, std::string token, int fromMenu, std::string serverID) {

    if (fromMenu == 1) {
        std::cout << "Please, enter the ID of the server you want to get the information about.\n";
        std::cout << "(Server ID is the numbers and letters after " + ip + "/server/<this is the ID>\n";
        std::cin >> serverID;

        if (serverID.length() != 8) {
            std::cout << "This doesn't look like a valid ID, the ID should be 8 characters long.\nThe ID you entered is " << serverID << "\n";
            std::cout << "Returning to main menu\n";
            menuReturn();
        } else {
            printConsole(ip, token, fromMenu, serverID);
        }
    }

    printConsole(ip, token, fromMenu, serverID);

}


#endif //PUFFERSTARTER_CLI_SEE_CONSOLE_H
