package controllers

import (
	"MemberRegisterApi/data"
	"MemberRegisterApi/model"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

var localDatabase data.LocalDatabase = data.NewLocalDatabase()

func StartWebServer() error {
	fmt.Println("Deploy REST API Server")
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", rootPage).Methods("GET")

	router.HandleFunc("/members", signUpMember).Methods("POST")
	router.HandleFunc("/members", fetchAllMembers).Methods("GET")

	router.HandleFunc("/members/{student_id}", fetchMemberById).Methods("GET")
	router.HandleFunc("/members/{student_id}", updateMember).Methods("PUT")
	router.HandleFunc("/members/{student_id}", deleteMember).Methods("DELETE")

	router.HandleFunc("/logs", takeLog).Methods("POST")
	router.HandleFunc("/logs", fetchAllLogs).Methods("GET")

	router.HandleFunc("/accesses", fetchAllAccesses).Methods("GET")
	router.HandleFunc("/accesses", registerAccess).Methods("POST")

	router.HandleFunc("/accesses/{student_id}", fetchAccessById).Methods("GET")
	router.HandleFunc("/accesses/{student_id}", updateAccess).Methods("PUT")

	return http.ListenAndServe(fmt.Sprintf(":%d", 8080), router)
}

func rootPage(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprint(w, "usage...")
	if err != nil {
		return
	}
}

func fetchAllMembers(w http.ResponseWriter, r *http.Request) {
	members := localDatabase.GetAllMembers()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(members)
}

func fetchMemberById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["student_id"]

	member := localDatabase.GetMemberById(key)
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(member)
}

func signUpMember(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var member model.Member
	if err := json.Unmarshal(reqBody, &member); err != nil {
		log.Fatal(err)
	}

	if _, err := localDatabase.SignUp(member); err != nil {
		fmt.Fprintf(w, "%s is already exists member", member.StudentId)

		return
	}

	memberList := localDatabase.GetAllMembers()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(memberList)
}

func updateMember(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["student_id"]

	var member model.Member

	reqBody, _ := ioutil.ReadAll(r.Body)
	if err := json.Unmarshal(reqBody, &member); err != nil {
		log.Fatal(err)
	}

	if key != member.StudentId {
		fmt.Fprintf(w, "path(%s) and body(%s) is not same", key, member.StudentId)
		return
	}

	if _, err := localDatabase.UpdateMember(member); err != nil {
		log.Fatal(err)
	}

	memberList := localDatabase.GetAllMembers()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(memberList)
}

func deleteMember(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["student_id"]

	if _, err := localDatabase.DeleteMember(key); err != nil {
		log.Fatal(err)
	}

	members := localDatabase.GetAllMembers()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(members)
}

func fetchAllLogs(w http.ResponseWriter, r *http.Request) {
	logs := localDatabase.GetAllLogs()

	if logs == nil {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(logs)
}

func takeLog(w http.ResponseWriter, r *http.Request) {
	var l model.Log
	reqBody, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(reqBody, &l)
	if err != nil {
		return
	}

	if result, err := localDatabase.TakeLog(l); err != nil {
		print(result, err.Error())
		return
	}
}

func registerAccess(w http.ResponseWriter, r *http.Request) {
	var access model.Access

	reqBody, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(reqBody, &access)
	if err != nil {
		return
	}

	localDatabase.RegisterAccess(access)
}

func updateAccess(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["student_id"]

	var access model.Access
	reqBody, _ := ioutil.ReadAll(r.Body)
	if err := json.Unmarshal(reqBody, &access); err != nil {
		return
	}

	if key != access.StudentId {
		fmt.Fprintf(w, "path(%s) and body(%s) is not same", key, access.StudentId)
		return
	}

	localDatabase.UpdateAccess(access)
}

func fetchAccessById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["student_id"]

	access := localDatabase.GetAccessById(key)

	json.NewEncoder(w).Encode(access)
}

func fetchAllAccesses(w http.ResponseWriter, r *http.Request) {
	accesses := localDatabase.GetAllAccess()

	w.Header().Set("Contents-Type", "application-json")
	json.NewEncoder(w).Encode(accesses)
}
