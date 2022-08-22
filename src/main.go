package main

import (
	"fmt"
	"github.com/tebeka/selenium"
	"net/http"
	"time"
)

func main() {

	http.HandleFunc("/bot", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("start check")
		// Connect to the WebDriver instance running locally.
		caps := selenium.Capabilities{"browserName": "firefox"}
		wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://127.0.0.1:4444/wd/hub"))
		if err != nil {
			panic(err)
		}
		fmt.Println("step 0")
		//defer func(wd selenium.WebDriver) {
		//	err = wd.Quit()
		//	if err != nil {
		//		panic(err)
		//	}
		//}(wd)

		err = wd.Get("https://www.instagram.com")

		if err != nil {
			panic(err)
		}
		fmt.Println("step 1")

		err = wd.SetImplicitWaitTimeout(time.Second * 10)

		fmt.Println("step 2")

		if err != nil {
			panic(err)
		}

		//formLogin, err := wd.FindElement(selenium.ByID, "loginForm")

		elementUserName, err := wd.FindElement(selenium.ByXPATH, "//input[@name='username']")
		if err != nil {
			panic(err)
		}

		err = elementUserName.SendKeys("__pinocchio__11")
		if err != nil {
			panic(err)
		}
		fmt.Println("fill username")

		elementPassword, err := wd.FindElement(selenium.ByXPATH, "//input[@name='password']")
		if err != nil {
			panic(err)
		}
		fmt.Println("fill password")
		err = elementPassword.SendKeys("Mp14111997")
		if err != nil {
			panic(err)
		}

		btnLogin, err := wd.FindElement(selenium.ByCSSSelector, "button[type=submit]")
		fmt.Println("click btn")
		err = btnLogin.Click()
		if err != nil {
			panic(err)
		}

		err = wd.Wait(waitForRedirect)
		if err != nil {
			panic(err)
		}

		//// skip save pass
		//skipSaveInfo, err := wd.FindElement(selenium.ByXPATH, "//*[@class='cmbtv']/button")
		//
		//if err != nil {
		//	panic(err)
		//}
		//
		//err = skipSaveInfo.Click()
		//if err != nil {
		//	panic(err)
		//}
		fmt.Println("redirect post")
		wd.Get("https://www.instagram.com/p/ChRnkiwJYsy/")

		err = wd.Wait(waitPostIsDisplay)
		if err != nil {
			panic(err)
		}

		commentElm, err := wd.FindElement(selenium.ByTagName, "textarea")

		fmt.Println("add comment")

		commentElm.Clear()
		err = commentElm.SendKeys("HI! I'M bot.....")
		if err != nil {
			panic(err)
		}

		fmt.Println("send comment")
		sendComment, err := wd.FindElement(selenium.ByXPATH, "//*[@type='submit']")
		err = sendComment.Click()

		if err != nil {
			panic(err)
		}

		_, err = w.Write([]byte("hello world1"))
		if err != nil {
			return
		}
	})

	err := http.ListenAndServe(":3001", nil)
	if err != nil {
		fmt.Println("server start at.")
		return
	}
	fmt.Println("server start at.")
}

var waitForRedirect = func(wd selenium.WebDriver) (bool, error) {

	url, err := wd.CurrentURL()
	if err != nil {

		return false, err
	}

	return url == "https://www.instagram.com/accounts/onetap/?next=%2F", nil

}

var waitPostIsDisplay = func(wd selenium.WebDriver) (bool, error) {

	commentElm, err := wd.FindElement(selenium.ByTagName, "textarea")
	if err != nil {
		return false, err
	}
	e, ok := commentElm.IsDisplayed()
	if ok != nil {
		panic(ok)
	}
	return e, nil

}
