// Package pubnubMessaging has the unit tests of package pubnubMessaging.
// pubnubDetailedHistory_test.go contains the tests related to the History requests on pubnub Api
package pubnubTests

import (
    "testing"
    "github.com/pubnub/go/3.4.1/pubnubMessaging"
    "strings"
    "fmt"
    "strconv"
    "encoding/json"
    "time"
)

// TestDetailedHistoryStart prints a message on the screen to mark the beginning of 
// detailed history tests.
// PrintTestMessage is defined in the common.go file.
func TestDetailedHistoryStart(t *testing.T){
    PrintTestMessage("==========DetailedHistory tests start==========")
}

// TestDetailedHistory publish's a message to a pubnub channel and when the sent response is received,
// calls the history method of the pubnubMessaging package to fetch 1 message. This received 
// message is compared to the message sent and if both match test is successful.  
func TestDetailedHistory(t *testing.T) {
    pubnubInstance := pubnubMessaging.PubnubInit("demo", "demo", "", "", false, "")    
    
    channel := "testChannel"
    message := "Test Message"

    returnPublishChannel := make(chan []byte)
    errorChannel := make(chan []byte)
    responseChannel := make(chan string)
    waitChannel := make(chan string)
    
    go pubnubInstance.Publish(channel, message, returnPublishChannel, errorChannel)
    go ParseResponse(returnPublishChannel, pubnubInstance, channel, message, "DetailedHistory", 1, responseChannel)
    go ParseErrorResponse(errorChannel, responseChannel)  
    go WaitForCompletion(responseChannel, waitChannel)
    ParseWaitResponse(waitChannel, t, "DetailedHistory")    
}

// TestEncryptedDetailedHistory publish's an encrypted message to a pubnub channel and when the 
// sent response is received, calls the history method of the pubnubMessaging package to fetch 
// 1 message. This received message is compared to the message sent and if both match test is successful.  
func TestEncryptedDetailedHistory(t *testing.T) {
    pubnubInstance := pubnubMessaging.PubnubInit("demo", "demo", "", "enigma", false, "")    
    
    channel := "testChannel"
    message := "Test Message"
    returnPublishChannel := make(chan []byte)
    errorChannel := make(chan []byte)
    responseChannel := make(chan string)
    waitChannel := make(chan string)

    go pubnubInstance.Publish(channel, message, returnPublishChannel, errorChannel)
    go ParseResponse(returnPublishChannel, pubnubInstance, channel, message, "EncryptedDetailedHistory", 1, responseChannel)
    go ParseErrorResponse(errorChannel, responseChannel)  
    go WaitForCompletion(responseChannel, waitChannel)
    ParseWaitResponse(waitChannel, t, "EncryptedDetailedHistory")
}

// TestDetailedHistoryFor10Messages publish's 10 unencrypted messages to a pubnub channel, and after that 
// calls the history method of the pubnubMessaging package to fetch last 10 messages. These received 
// messages are compared to the messages sent and if all match test is successful.  
func TestDetailedHistoryFor10Messages(t *testing.T) {
    testName := "TestDetailedHistoryFor10Messages"
    DetailedHistoryFor10Messages(t, "", testName)
    time.Sleep(2 * time.Second)
}

// TestDetailedHistoryFor10EncryptedMessages publish's 10 encrypted messages to a pubnub channel, and after that 
// calls the history method of the pubnubMessaging package to fetch last 10 messages. These received 
// messages are compared to the messages sent and if all match test is successful.  
func TestDetailedHistoryFor10EncryptedMessages(t *testing.T) {
    testName := "TestDetailedHistoryFor10EncryptedMessages"
    DetailedHistoryFor10Messages(t, "enigma", testName)
    time.Sleep(2 * time.Second)
}

// DetailedHistoryFor10Messages is a common method used by both TestDetailedHistoryFor10EncryptedMessages 
// and TestDetailedHistoryFor10Messages to publish's 10 messages to a pubnub channel, and after that 
// call the history method of the pubnubMessaging package to fetch last 10 messages. These received 
// messages are compared to the messages sent and if all match test is successful.  
func DetailedHistoryFor10Messages(t *testing.T, cipherKey string, testName string) {
    numberOfMessages := 10
    startMessagesFrom := 0
    pubnubInstance := pubnubMessaging.PubnubInit("demo", "demo", "", cipherKey, false, "")    
    
    message := "Test Message "
    channel := "testChannel"
    
    messagesSent := PublishMessages(pubnubInstance, channel, t, startMessagesFrom, numberOfMessages, message)
    
    if(messagesSent){
        returnHistoryChannel := make(chan []byte)
        errorChannel := make(chan []byte)
        responseChannel := make(chan string)
        waitChannel := make(chan string)
            
        go pubnubInstance.History(channel, numberOfMessages, 0, 0, false, returnHistoryChannel, errorChannel)
        go ParseHistoryResponseForMultipleMessages(returnHistoryChannel, channel, message, testName, startMessagesFrom, numberOfMessages, cipherKey, responseChannel)
        go ParseErrorResponse(errorChannel, responseChannel)  
        go WaitForCompletion(responseChannel, waitChannel)
        ParseWaitResponse(waitChannel, t, testName)
    }else{
        t.Error("Test '" + testName + "': failed.");
    }    
}

// TestDetailedHistoryParamsFor10MessagesWithSeretKey publish's 10 unencrypted secret keyed messages 
// to a pubnub channel, and after that calls the history method of the pubnubMessaging package to fetch 
// last 10 messages with time parameters between which the messages were sent. These received 
// messages are compared to the messages sent and if all match test is successful.  
func TestDetailedHistoryParamsFor10MessagesWithSeretKey(t *testing.T) {
    testName := "TestDetailedHistoryFor10MessagesWithSeretKey"
    DetailedHistoryParamsFor10Messages(t, "", "secret", testName)
    time.Sleep(2 * time.Second)
}

// TestDetailedHistoryParamsFor10EncryptedMessagesWithSeretKey publish's 10 encrypted secret keyed messages 
// to a pubnub channel, and after that calls the history method of the pubnubMessaging package to fetch 
// last 10 messages with time parameters between which the messages were sent. These received 
// messages are compared to the messages sent and if all match test is successful.  
func TestDetailedHistoryParamsFor10EncryptedMessagesWithSeretKey(t *testing.T) {
    testName := "TestDetailedHistoryFor10EncryptedMessagesWithSeretKey"
    DetailedHistoryParamsFor10Messages(t, "enigma", "secret", testName)
    time.Sleep(2 * time.Second)
}

// TestDetailedHistoryParamsFor10Messages publish's 10 unencrypted messages 
// to a pubnub channel, and after that calls the history method of the pubnubMessaging package to fetch 
// last 10 messages with time parameters between which the messages were sent. These received 
// messages are compared to the messages sent and if all match test is successful.  
func TestDetailedHistoryParamsFor10Messages(t *testing.T) {
    testName := "TestDetailedHistoryFor10Messages"
    DetailedHistoryParamsFor10Messages(t, "", "", testName)
    time.Sleep(2 * time.Second)
}

// TestDetailedHistoryParamsFor10EncryptedMessages publish's 10 encrypted messages 
// to a pubnub channel, and after that calls the history method of the pubnubMessaging package to fetch 
// last 10 messages with time parameters between which the messages were sent. These received 
// messages are compared to the messages sent and if all match test is successful.  
func TestDetailedHistoryParamsFor10EncryptedMessages(t *testing.T) {
    testName := "TestDetailedHistoryParamsFor10EncryptedMessages"
    DetailedHistoryParamsFor10Messages(t, "enigma", "", testName)
    time.Sleep(2 * time.Second)
}

// DetailedHistoryFor10Messages is a common method used by both TestDetailedHistoryFor10EncryptedMessages 
// and TestDetailedHistoryFor10Messages to publish's 10 messages to a pubnub channel, and after that 
// call the history method of the pubnubMessaging package to fetch last 10 messages with time parameters 
// between which the messages were sent. These received message is compared to the messages sent and 
// if all match test is successful.  
func DetailedHistoryParamsFor10Messages(t *testing.T, cipherKey string, secretKey string, testName string) {
    numberOfMessages := 5
    pubnubInstance := pubnubMessaging.PubnubInit("demo", "demo", secretKey, cipherKey, false, "")    
    
    message := "Test Message "
    channel := "testChannel"
    
    startTime := GetServerTime(pubnubInstance, t, testName) 
    startMessagesFrom := 0
    messagesSent := PublishMessages(pubnubInstance, channel, t, startMessagesFrom, numberOfMessages, message)
    
    midTime := GetServerTime(pubnubInstance, t, testName)
    startMessagesFrom = 5
    messagesSent2 := PublishMessages(pubnubInstance, channel, t, startMessagesFrom, numberOfMessages, message)
    endTime := GetServerTime(pubnubInstance, t, testName)
    
    startMessagesFrom = 0
    if(messagesSent){
        returnHistoryChannel := make(chan []byte)
        responseChannel := make(chan string)
        errorChannel := make(chan []byte)
        waitChannel := make(chan string)
            
        go pubnubInstance.History(channel, numberOfMessages, startTime, midTime, false, returnHistoryChannel, errorChannel)
        go ParseHistoryResponseForMultipleMessages(returnHistoryChannel, channel, message, testName, startMessagesFrom, numberOfMessages, cipherKey, responseChannel)
        go ParseErrorResponse(errorChannel, responseChannel)  
        go WaitForCompletion(responseChannel, waitChannel)
        ParseWaitResponse(waitChannel, t, testName)
    }else{
        t.Error("Test '" + testName + "': failed.");
    }
    
    startMessagesFrom = 5
    if(messagesSent2){
        returnHistoryChannel2 := make(chan []byte)
        errorChannel2 := make(chan []byte)
        responseChannel2 := make(chan string)
        waitChannel2 := make(chan string)
            
        go pubnubInstance.History(channel, numberOfMessages, midTime, endTime, false, returnHistoryChannel2, errorChannel2)
        go ParseHistoryResponseForMultipleMessages(returnHistoryChannel2, channel, message, testName, startMessagesFrom, numberOfMessages, cipherKey, responseChannel2)
        go ParseErrorResponse(errorChannel2, responseChannel2)  
        go WaitForCompletion(responseChannel2, waitChannel2)
        ParseWaitResponse(waitChannel2, t, testName)
    }else{
        t.Error("Test '" + testName + "': failed.");
    }    
}

// GetServerTime calls the GetTime method of the pubnubMessaging, parses the response to get the
// value and return it.
func GetServerTime(pubnubInstance *pubnubMessaging.Pubnub, t *testing.T, testName string) int64{
    returnTimeChannel := make(chan []byte)
    errorChannel := make(chan []byte)
    go pubnubInstance.GetTime(returnTimeChannel, errorChannel)
    return ParseServerTimeResponse(returnTimeChannel, t, testName)    
}

// ParseServerTimeResponse unmarshals the time response from the pubnub api and returns the int64 value.
// On error the test fails.
func ParseServerTimeResponse(returnChannel chan []byte,t *testing.T, testName string) int64 {
    for {
        value, ok := <-returnChannel
        if !ok {
            break
        }
        if string(value) != "[]"{
            response := string(value)
            if(response != ""){
                var arr []int64
                err2 := json.Unmarshal(value, &arr)
                if(err2 != nil){
                    fmt.Println("err2 time", err2)
                    t.Error("Test '" + testName + "': failed.");
                    break
                }else {    
                    return arr[0]
                }         
            } else {
                fmt.Println("response", response)
                t.Error("Test '" + testName + "': failed.");
                break
            }
        }
    }
    return 0
}

// PublishMessages calls the publish method of pubnubMessaging package numberOfMessages times 
// and appends the count with the message to distinguish from the others.
// 
// Parameters:
// pubnubInstance: a reference of *pubnubMessaging.Pubnub, 
// channel: the pubnub channel to publish the messages, 
// t: a reference to *testing.T,
// startMessagesFrom: the message identifer, 
// numberOfMessages: number of messages to send,
// message: message to send.
// 
// returns a bool if the publish of all messages is successful.
func PublishMessages(pubnubInstance *pubnubMessaging.Pubnub, channel string, t *testing.T, startMessagesFrom int, numberOfMessages int, message string) bool{
    messagesReceived := 0
    messageToSend := ""
    for i:=startMessagesFrom; i< startMessagesFrom+numberOfMessages; i++{
        messageToSend = message + strconv.Itoa(i)
    
        returnPublishChannel := make(chan []byte)
        errorChannel := make(chan []byte)
        go pubnubInstance.Publish(channel, messageToSend, returnPublishChannel, errorChannel)
        messagesReceived++
        time.Sleep(500*time.Millisecond)
    }
    if(messagesReceived == numberOfMessages){
        return true
    }
    
    return false
}

// ParseHistoryResponseForMultipleMessages unmarshalls the response of the history call to the
// pubnub api and compares the received messages to the sent messages. If the response match the 
// test is successful.
// 
// Parameters:
// returnChannel: channel to read the response from,
// t: a reference to *testing.T, 
// channel: the pubnub channel to publish the messages,
// message: message to compare,
// testname: the test name form where this method is called,
// startMessagesFrom: the message identifer, 
// numberOfMessages: number of messages to send,
// cipherKey: the cipher key if used. Can be empty.
func ParseHistoryResponseForMultipleMessages(returnChannel chan []byte, channel string, message string, testName string, startMessagesFrom int, numberOfMessages int, cipherKey string, responseChannel chan string){
    for {
        value, ok := <-returnChannel
        if !ok {
            break
        }
        if string(value) != "[]"{
            data, _, _, err := pubnubMessaging.ParseJson(value, cipherKey)
            if(err != nil) {
                //t.Error("Test '" + testName + "': failed.")
                responseChannel <- "Test '" + testName + "': failed. Message: " + err.Error()
            } else {
                var arr []string
                err2 := json.Unmarshal([]byte(data), &arr)
                if(err2 != nil){
                    //t.Error("Test '" + testName + "': failed.");
                    responseChannel <- "Test '" + testName + "': failed. Message: " + err2.Error()
                }else {    
                    messagesReceived := 0
                    
                    if(len(arr) != numberOfMessages){
                        responseChannel <- "Test '" + testName + "': failed."
                        //t.Error("Test '" + testName + "': failed.");
                        break
                    }
                    for i:=0; i < numberOfMessages; i++ {
                        if(arr[i] == message + strconv.Itoa(startMessagesFrom+i)){
                            //fmt.Println("data:",arr[i])
                            messagesReceived++
                        }
                    }   
                    if(messagesReceived == numberOfMessages){
                        fmt.Println("Test '" + testName + "': passed.")
                        responseChannel <- "Test '" + testName + "': passed."
                    } else {
                        responseChannel <- "Test '" + testName + "': failed. Returned message mismatch"
                        //t.Error("Test '" + testName + "': failed.");
                    }
                break
                }
            }
        }
    }
}

// ParseHistoryResponse parses the history response from the pubnub api on the returnChannel
// and checks if the response contains the message. If true then the test is successful. 
func ParseHistoryResponse(returnChannel chan []byte, channel string, message string, testName string, responseChannel chan string){
    for {
        value, ok := <-returnChannel
        if !ok {
            break
        }
        if string(value) != "[]"{
            response := string(value)
            //fmt.Println("response", response)
            
            if(strings.Contains(response, message)){
                //fmt.Println("Test '" + testName + "': passed.")
                responseChannel <- "Test '" + testName + "': passed." 
                break
            } else {
                responseChannel <- "Test '" + testName + "': failed."                
            }
        }
    }
}

// ParseResponse parses the publish response from the pubnub api on the returnChannel and 
// when the sent response is received, calls the history method of the pubnubMessaging 
// package to fetch 1 message. 
func ParseResponse(returnChannel chan []byte, pubnubInstance *pubnubMessaging.Pubnub, channel string, message string, testName string, numberOfMessages int, responseChannel chan string){
    for {
        value, ok := <-returnChannel
        if !ok {
            break
        }
        if string(value) != "[]"{
            returnHistoryChannel := make(chan []byte)
            var errorChannel = make(chan []byte)
            go pubnubInstance.History(channel, 1, 0, 0, false, returnHistoryChannel, errorChannel)
            go ParseHistoryResponse(returnHistoryChannel, channel, message, testName, responseChannel)
            go ParseErrorResponse(errorChannel, responseChannel)  
            break
        }
    }
}

// TestDetailedHistoryEnd prints a message on the screen to mark the end of 
// detailed history tests.
// PrintTestMessage is defined in the common.go file.
func TestDetailedHistoryEnd(t *testing.T){
    PrintTestMessage("==========DetailedHistory tests end==========")
}