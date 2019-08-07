package preparer

import (
	"encoding/base64"
	"fmt"
	"testing"

	"github.com/Cepreu/gofrend/ivr"
)

var (
	user, pwd = "sam2@007.com", "pwd1234567"
	auth      = base64.StdEncoding.EncodeToString([]byte(user + ":" + pwd))
	campaign  = "sergei_inbound"
)

var script = ivr.IVRScript{Domain: "qwerty", Name: "rememberme",
	Modules:   make(map[ivr.ModuleID]ivr.Module),
	Variables: make(map[ivr.VariableID]*ivr.Variable),
}

func TestGetIvrFromF9(t *testing.T) {
	ivr, err := getIvrFromF9(auth, "delme")
	if err != nil || len(ivr) == 0 {
		t.Errorf("TestGetIvrFromF9: \n%v", err)
	}
}

func TestCreateIvrF9(t *testing.T) {
	generatedScriptName := script.Name + "_generated"

	resp, err := queryF9(auth, func() string { return createIVRscriptContent(generatedScriptName) })
	if err != nil {
		fault, _ := getf9errorDescr(resp)
		exp := fmt.Sprintf("IvrScript with name \"%s\" already exists", generatedScriptName)
		if exp != fault.FaultString {
			t.Errorf("TestCreateIvrF9: expected OK or \"%s\", received \"%s\"", exp, fault.FaultString)
		}
	}
}

func TestConfigureF9(t *testing.T) {
	err := configureF9(auth, campaign, &script)
	if err != nil {
		t.Errorf("TestConfigureF9: \n%v", err)
	}
}

func TestChangePwdF9(t *testing.T) {
	auth := base64.StdEncoding.EncodeToString([]byte(user + ":" + pwd))
	err, pwd := changeUserPwdF9(auth)
	//	if err != nil {
	t.Errorf("TestChangePwdF9: \n%v, pwd=%s", err, pwd)
	//	}
}

func TestGenerateIVRContent(t *testing.T) {

	val, _ := ivr.NewTimeValue(60)
	script.Variables["var_time"] = ivr.NewVariable("var_time", "time var", ivr.ValTime, val)
	val, _ = ivr.NewDateValue(2019, 6, 7)
	script.Variables["var_date"] = ivr.NewVariable("var_date", "date var", ivr.ValDate, val)
	val, _ = ivr.NewIntegerValue(120)
	script.Variables["var_int"] = ivr.NewVariable("var_int", "integer var", ivr.ValInteger, val)
	val, _ = ivr.NewNumericValue(3.14)
	script.Variables["var_num"] = ivr.NewVariable("var_num", "numeric var", ivr.ValNumeric, val)
	val, _ = ivr.NewUSCurrencyValue(100.50)
	script.Variables["var_currency"] = ivr.NewVariable("var_currency", "currency evro", ivr.ValCurrencyEuro, val)
	val, _ = ivr.NewKeyValue(`{"a":"sdfg","q":"werty"}`)
	script.Variables["var_kvlist"] = ivr.NewVariable("var_kvlist", "KV list", ivr.ValKVList, val)

	var hu1 = ivr.HangupModule{ErrCode: "duration", OverwriteDisp: true}
	hu1.SetGeneralInfo("Hangup1", "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA1",
		[]ivr.ModuleID{"ED132095BE1E4F47B51DA0BB842C3EEF", "F1E142D8CF27471D8940713A637A1C1D"},
		"", "", "No Disposition", "false")
	script.Modules["Hangup1"] = &hu1

	var hu2 = ivr.HangupModule{ErrCode: "attitude", OverwriteDisp: false}
	hu2.SetGeneralInfo("Hangup2", "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA2",
		[]ivr.ModuleID{"TTTTTTTTTTTTTTTTTTTTTTTTTTTTTTT1"},
		"", "", "No Disposition", "false")
	script.Modules["Hangup2"] = &hu2

	var st1 = ivr.SkillTransferModule{
		Data: struct {
			InnerXML string `xml:",innerxml" datastore:",noindex"`
		}{
			InnerXML: "\n\t\t\u003cdispo\u003e\n\t\t\t\u003cid\u003e-5\u003c/id\u003e\n\t\t\t\u003cname\u003eAbandon\u003c/name\u003e\n\t\t\u003c/dispo\u003e\n\t\t\u003cmaxQueueTime\u003e600\u003c/maxQueueTime\u003e\n\t\t\u003cqueueIfOnCall\u003etrue\u003c/queueIfOnCall\u003e\n\t\t\u003conCallQueueTime\u003e601\u003c/onCallQueueTime\u003e\n\t\t\u003cqueueIfOnBreakOrLoggedOut\u003etrue\u003c/queueIfOnBreakOrLoggedOut\u003e\n\t\t\u003conBreakOrLoggedOutQueueTime\u003e602\u003c/onBreakOrLoggedOutQueueTime\u003e\n\t\t\u003cterminationDigit\u003e35\u003c/terminationDigit\u003e\n\t\t\u003ccallbackDigit\u003e42\u003c/callbackDigit\u003e\n\t\t\u003conQueueTimeoutExpiration\u003etrue\u003c/onQueueTimeoutExpiration\u003e\n\t\t\u003cpauseBeforeTransfer\u003e0\u003c/pauseBeforeTransfer\u003e\n\t\t\u003cmaxRingTime\u003e15\u003c/maxRingTime\u003e\n\t\t\u003cplaceOnBreakIfNoAnswer\u003etrue\u003c/placeOnBreakIfNoAnswer\u003e\n\t\t\u003cvmTransferOnQueueTimeout\u003efalse\u003c/vmTransferOnQueueTimeout\u003e\n\t\t\u003cvmTransferOnDigit\u003etrue\u003c/vmTransferOnDigit\u003e\n\t\t\u003cvmTransferDigit\u003e51\u003c/vmTransferDigit\u003e\n\t\t\u003cvmBoxType\u003eSKILL\u003c/vmBoxType\u003e\n\t\t\u003cvmSkillBox\u003e\n\t\t\t\u003cid\u003e266194\u003c/id\u003e\n\t\t\t\u003cname\u003esk_skill1\u003c/name\u003e\n\t\t\u003c/vmSkillBox\u003e\n\t\t\u003cclearDigitBuffer\u003etrue\u003c/clearDigitBuffer\u003e\n\t\t\u003cenableMusicOnHold\u003etrue\u003c/enableMusicOnHold\u003e\n\t\t\u003cannouncements\u003e\n\t\t\t\u003cenabled\u003etrue\u003c/enabled\u003e\n\t\t\t\u003cloopped\u003etrue\u003c/loopped\u003e\n\t\t\t\u003ctimeout\u003e3\u003c/timeout\u003e\n\t\t\t\u003cprompt/\u003e\n\t\t\t\u003cannType\u003eEWT\u003c/annType\u003e\n\t\t\u003c/announcements\u003e\n\t\t\u003cannouncements\u003e\n\t\t\t\u003cenabled\u003etrue\u003c/enabled\u003e\n\t\t\t\u003cloopped\u003efalse\u003c/loopped\u003e\n\t\t\t\u003ctimeout\u003e0\u003c/timeout\u003e\n\t\t\t\u003cprompt\u003e\n\t\t\t\t\u003cid\u003e222995\u003c/id\u003e\n\t\t\t\t\u003cname\u003ehold music\u003c/name\u003e\n\t\t\t\u003c/prompt\u003e\n\t\t\t\u003cannType\u003ePROMPT\u003c/annType\u003e\n\t\t\u003c/announcements\u003e\n\t\t\u003cannouncements\u003e\n\t\t\t\u003cenabled\u003efalse\u003c/enabled\u003e\n\t\t\t\u003cloopped\u003efalse\u003c/loopped\u003e\n\t\t\t\u003ctimeout\u003e1\u003c/timeout\u003e\n\t\t\t\u003cprompt/\u003e\n\t\t\t\u003cannType\u003eEWT\u003c/annType\u003e\n\t\t\u003c/announcements\u003e\n\t\t\u003cannouncements\u003e\n\t\t\t\u003cenabled\u003efalse\u003c/enabled\u003e\n\t\t\t\u003cloopped\u003efalse\u003c/loopped\u003e\n\t\t\t\u003ctimeout\u003e1\u003c/timeout\u003e\n\t\t\t\u003cprompt/\u003e\n\t\t\t\u003cannType\u003eEWT\u003c/annType\u003e\n\t\t\u003c/announcements\u003e\n\t\t\u003cannouncements\u003e\n\t\t\t\u003cenabled\u003efalse\u003c/enabled\u003e\n\t\t\t\u003cloopped\u003efalse\u003c/loopped\u003e\n\t\t\t\u003ctimeout\u003e1\u003c/timeout\u003e\n\t\t\t\u003cprompt/\u003e\n\t\t\t\u003cannType\u003eEWT\u003c/annType\u003e\n\t\t\u003c/announcements\u003e\n\t\t\u003cconnectors\u003e\n\t\t\t\u003cid\u003e119033\u003c/id\u003e\n\t\t\t\u003cname\u003eConn4_Accept\u003c/name\u003e\n\t\t\u003c/connectors\u003e\n\t\t\u003cpriorityChangeType\u003eINCREASE\u003c/priorityChangeType\u003e\n\t\t\u003cpriorityChangeValue\u003e\n\t\t\t\u003cisVarSelected\u003efalse\u003c/isVarSelected\u003e\n\t\t\t\u003cintegerValue\u003e\n\t\t\t\t\u003cvalue\u003e10\u003c/value\u003e\n\t\t\t\u003c/integerValue\u003e\n\t\t\u003c/priorityChangeValue\u003e\n\t\t\u003clistOfSkillsEx\u003e\n\t\t\t\u003cextrnalObj\u003e\n\t\t\t\t\u003cid\u003e266194\u003c/id\u003e\n\t\t\t\t\u003cname\u003esk_skill1\u003c/name\u003e\n\t\t\t\u003c/extrnalObj\u003e\n\t\t\t\u003cvarSelected\u003efalse\u003c/varSelected\u003e\n\t\t\u003c/listOfSkillsEx\u003e\n\t\t\u003clistOfSkillsEx\u003e\n\t\t\t\u003cextrnalObj\u003e\n\t\t\t\t\u003cid\u003e266254\u003c/id\u003e\n\t\t\t\t\u003cname\u003emk_skill1\u003c/name\u003e\n\t\t\t\u003c/extrnalObj\u003e\n\t\t\t\u003cvarSelected\u003efalse\u003c/varSelected\u003e\n\t\t\u003c/listOfSkillsEx\u003e\n\t\t\u003ctaskType\u003e0\u003c/taskType\u003e\n\t\t\u003ctransferAlgorithm\u003e\n\t\t\t\u003calgorithmType\u003eLongestReadyTime\u003c/algorithmType\u003e\n\t\t\t\u003cstatAlgorithmTimeWindow\u003eE15minutes\u003c/statAlgorithmTimeWindow\u003e\n\t\t\u003c/transferAlgorithm\u003e\n\t\t\u003crecordedFilesAction\u003eKEEP_AS_RECORDINGD\u003c/recordedFilesAction\u003e\n\t\t\u003ccallbackNumberToCav\u003e\n\t\t\t\u003cid\u003e201\u003c/id\u003e\n\t\t\t\u003cname\u003eqwerty\u003c/name\u003e\n\t\t\u003c/callbackNumberToCav\u003e\n\t\t\u003ccallbackNumberFromCav\u003e\n\t\t\t\u003cid\u003e63\u003c/id\u003e\n\t\t\t\u003cname\u003enumber\u003c/name\u003e\n\t\t\u003c/callbackNumberFromCav\u003e\n\t\t\u003ccallbackPhoneNumberPrompt\u003e\n\t\t\t\u003cid\u003e-26\u003c/id\u003e\n\t\t\t\u003cname\u003eQC1YourPhoneNumberIs\u003c/name\u003e\n\t\t\u003c/callbackPhoneNumberPrompt\u003e\n\t\t\u003ccallbackConfirmingPhoneNumberPrompt\u003e\n\t\t\t\u003cid\u003e-27\u003c/id\u003e\n\t\t\t\u003cname\u003eQC2PressOneIfCorrect\u003c/name\u003e\n\t\t\u003c/callbackConfirmingPhoneNumberPrompt\u003e\n\t\t\u003ccallbackEnteringPhoneNumberPrompt\u003e\n\t\t\t\u003cid\u003e-28\u003c/id\u003e\n\t\t\t\u003cname\u003eQC3PleaseEnterYourPhoneNumber\u003c/name\u003e\n\t\t\u003c/callbackEnteringPhoneNumberPrompt\u003e\n\t\t\u003ccallbackRecordingCallerNamePrompt\u003e\n\t\t\t\u003cid\u003e-36\u003c/id\u003e\n\t\t\t\u003cname\u003eQC5PleaseClearlySayYourName\u003c/name\u003e\n\t\t\u003c/callbackRecordingCallerNamePrompt\u003e\n\t\t\u003ccallbackConfirmationPrompt\u003e\n\t\t\t\u003cid\u003e-29\u003c/id\u003e\n\t\t\t\u003cname\u003eQC4YouWillBeCalled\u003c/name\u003e\n\t\t\u003c/callbackConfirmationPrompt\u003e\n\t\t\u003ccallbackQueueTimeoutSec\u003e18000\u003c/callbackQueueTimeoutSec\u003e\n\t\t\u003ccallbackEnterDigitsMaxTimeSec\u003e5\u003c/callbackEnterDigitsMaxTimeSec\u003e\n\t\t\u003ccallbackAllowInternational\u003efalse\u003c/callbackAllowInternational\u003e\n\t\t\u003cisTcpaConsentEnabled\u003etrue\u003c/isTcpaConsentEnabled\u003e\n\t\t\u003ctcpaConsentText\u003eH4sIAAAAAAAAADWOwQ3DMAwDV+EARYboL58uUPSh2EwjwJYCy878TVDkS/IOnJHcgtbRHY2JehBJ\nSgmszSvezxFqjMBLKj+Qc7gRnYX75kbYqAsb9uaHZmbI4gcfUEtlZLXv1VTv6iblL55mDMts0cXy\naTuV9wcNmHfIFWS9IPiKfbS0SXD6AQLT2xmtAAAA\u003c/tcpaConsentText\u003e\n\t\t\u003cignoreSkillsOrder\u003efalse\u003c/ignoreSkillsOrder\u003e\n\t\t\u003crecordCallerNameOnQueueCallback\u003efalse\u003c/recordCallerNameOnQueueCallback\u003e\n\t",
		},
	}
	st1.SetGeneralInfo("SkillTransfer1", "TTTTTTTTTTTTTTTTTTTTTTTTTTTTTTT1",
		[]ivr.ModuleID{"E97CB11F6641449FB9D34E6A1C38326C"}, "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA2", "",
		"", "false")
	script.Modules["SkillTransfer1"] = &st1

	res, err := generateIVRContent(&script)
	fmt.Println(res)
	if err != nil {
		t.Errorf("TestGenerateIVRContent: Error: %v", err)
	}
}
