package xmlparser

import (
	"encoding/json"
	"encoding/xml"
	"strings"
	"testing"

	"github.com/Cepreu/gofrend/ivr"
	"golang.org/x/net/html/charset"
)

func TestSkillTransfer(t *testing.T) {
	var xmlData = `<skillTransfer>
	<ascendants>E97CB11F6641449FB9D34E6A1C38326C</ascendants>
	<singleDescendant>E493C28EA263451CBECAA007C416F4FF</singleDescendant>
	<moduleName>SkillTransfer4</moduleName>
	<locationX>254</locationX>
	<locationY>107</locationY>
	<moduleId>39FE8C2F1E2744D6975EF11767639A80</moduleId>
	<data>
		<dispo>
			<id>-5</id>
			<name>Abandon</name>
		</dispo>
		<maxQueueTime>600</maxQueueTime>
		<queueIfOnCall>true</queueIfOnCall>
		<onCallQueueTime>601</onCallQueueTime>
		<queueIfOnBreakOrLoggedOut>true</queueIfOnBreakOrLoggedOut>
		<onBreakOrLoggedOutQueueTime>602</onBreakOrLoggedOutQueueTime>
		<terminationDigit>35</terminationDigit>
		<callbackDigit>42</callbackDigit>
		<onQueueTimeoutExpiration>true</onQueueTimeoutExpiration>
		<pauseBeforeTransfer>0</pauseBeforeTransfer>
		<maxRingTime>15</maxRingTime>
		<placeOnBreakIfNoAnswer>true</placeOnBreakIfNoAnswer>
		<vmTransferOnQueueTimeout>false</vmTransferOnQueueTimeout>
		<vmTransferOnDigit>true</vmTransferOnDigit>
		<vmTransferDigit>51</vmTransferDigit>
		<vmBoxType>SKILL</vmBoxType>
		<vmSkillBox>
			<id>266194</id>
			<name>sk_skill1</name>
		</vmSkillBox>
		<clearDigitBuffer>true</clearDigitBuffer>
		<enableMusicOnHold>true</enableMusicOnHold>
		<announcements>
			<enabled>true</enabled>
			<loopped>true</loopped>
			<timeout>3</timeout>
			<prompt/>
			<annType>EWT</annType>
		</announcements>
		<announcements>
			<enabled>true</enabled>
			<loopped>false</loopped>
			<timeout>0</timeout>
			<prompt>
				<id>222995</id>
				<name>hold music</name>
			</prompt>
			<annType>PROMPT</annType>
		</announcements>
		<announcements>
			<enabled>false</enabled>
			<loopped>false</loopped>
			<timeout>1</timeout>
			<prompt/>
			<annType>EWT</annType>
		</announcements>
		<announcements>
			<enabled>false</enabled>
			<loopped>false</loopped>
			<timeout>1</timeout>
			<prompt/>
			<annType>EWT</annType>
		</announcements>
		<announcements>
			<enabled>false</enabled>
			<loopped>false</loopped>
			<timeout>1</timeout>
			<prompt/>
			<annType>EWT</annType>
		</announcements>
		<connectors>
			<id>119033</id>
			<name>Conn4_Accept</name>
		</connectors>
		<priorityChangeType>INCREASE</priorityChangeType>
		<priorityChangeValue>
			<isVarSelected>false</isVarSelected>
			<integerValue>
				<value>10</value>
			</integerValue>
		</priorityChangeValue>
		<listOfSkillsEx>
			<extrnalObj>
				<id>266194</id>
				<name>sk_skill1</name>
			</extrnalObj>
			<varSelected>false</varSelected>
		</listOfSkillsEx>
		<listOfSkillsEx>
			<extrnalObj>
				<id>266254</id>
				<name>mk_skill1</name>
			</extrnalObj>
			<varSelected>false</varSelected>
		</listOfSkillsEx>
		<taskType>0</taskType>
		<transferAlgorithm>
			<algorithmType>LongestReadyTime</algorithmType>
			<statAlgorithmTimeWindow>E15minutes</statAlgorithmTimeWindow>
		</transferAlgorithm>
		<recordedFilesAction>KEEP_AS_RECORDINGD</recordedFilesAction>
		<callbackNumberToCav>
			<id>201</id>
			<name>qwerty</name>
		</callbackNumberToCav>
		<callbackNumberFromCav>
			<id>63</id>
			<name>number</name>
		</callbackNumberFromCav>
		<callbackPhoneNumberPrompt>
			<id>-26</id>
			<name>QC1YourPhoneNumberIs</name>
		</callbackPhoneNumberPrompt>
		<callbackConfirmingPhoneNumberPrompt>
			<id>-27</id>
			<name>QC2PressOneIfCorrect</name>
		</callbackConfirmingPhoneNumberPrompt>
		<callbackEnteringPhoneNumberPrompt>
			<id>-28</id>
			<name>QC3PleaseEnterYourPhoneNumber</name>
		</callbackEnteringPhoneNumberPrompt>
		<callbackRecordingCallerNamePrompt>
			<id>-36</id>
			<name>QC5PleaseClearlySayYourName</name>
		</callbackRecordingCallerNamePrompt>
		<callbackConfirmationPrompt>
			<id>-29</id>
			<name>QC4YouWillBeCalled</name>
		</callbackConfirmationPrompt>
		<callbackQueueTimeoutSec>18000</callbackQueueTimeoutSec>
		<callbackEnterDigitsMaxTimeSec>5</callbackEnterDigitsMaxTimeSec>
		<callbackAllowInternational>false</callbackAllowInternational>
		<isTcpaConsentEnabled>true</isTcpaConsentEnabled>
		<tcpaConsentText>H4sIAAAAAAAAADWOwQ3DMAwDV+EARYboL58uUPSh2EwjwJYCy878TVDkS/IOnJHcgtbRHY2JehBJ
SgmszSvezxFqjMBLKj+Qc7gRnYX75kbYqAsb9uaHZmbI4gcfUEtlZLXv1VTv6iblL55mDMts0cXy
aTuV9wcNmHfIFWS9IPiKfbS0SXD6AQLT2xmtAAAA</tcpaConsentText>
		<ignoreSkillsOrder>false</ignoreSkillsOrder>
		<recordCallerNameOnQueueCallback>false</recordCallerNameOnQueueCallback>
	</data>
</skillTransfer>`

	decoder := xml.NewDecoder(strings.NewReader(xmlData))
	decoder.CharsetReader = charset.NewReaderLabel
	_, _ = decoder.Token()

	s := &ivr.IVRScript{
		Variables: make(ivr.Variables),
	}
	res := newSkillTransferModule(decoder, s)
	if res == nil {
		t.Fatal("nSkill Transfer module wasn't parsed...")
	}
	var mhu = (res.(xmlSkillTransferModule)).m

	var expected = ivr.SkillTransferModule{
		Data: struct {
			InnerXML string `xml:",innerxml" datastore:",noindex"`
		}{
			InnerXML: "\n\t\t\u003cdispo\u003e\n\t\t\t\u003cid\u003e-5\u003c/id\u003e\n\t\t\t\u003cname\u003eAbandon\u003c/name\u003e\n\t\t\u003c/dispo\u003e\n\t\t\u003cmaxQueueTime\u003e600\u003c/maxQueueTime\u003e\n\t\t\u003cqueueIfOnCall\u003etrue\u003c/queueIfOnCall\u003e\n\t\t\u003conCallQueueTime\u003e601\u003c/onCallQueueTime\u003e\n\t\t\u003cqueueIfOnBreakOrLoggedOut\u003etrue\u003c/queueIfOnBreakOrLoggedOut\u003e\n\t\t\u003conBreakOrLoggedOutQueueTime\u003e602\u003c/onBreakOrLoggedOutQueueTime\u003e\n\t\t\u003cterminationDigit\u003e35\u003c/terminationDigit\u003e\n\t\t\u003ccallbackDigit\u003e42\u003c/callbackDigit\u003e\n\t\t\u003conQueueTimeoutExpiration\u003etrue\u003c/onQueueTimeoutExpiration\u003e\n\t\t\u003cpauseBeforeTransfer\u003e0\u003c/pauseBeforeTransfer\u003e\n\t\t\u003cmaxRingTime\u003e15\u003c/maxRingTime\u003e\n\t\t\u003cplaceOnBreakIfNoAnswer\u003etrue\u003c/placeOnBreakIfNoAnswer\u003e\n\t\t\u003cvmTransferOnQueueTimeout\u003efalse\u003c/vmTransferOnQueueTimeout\u003e\n\t\t\u003cvmTransferOnDigit\u003etrue\u003c/vmTransferOnDigit\u003e\n\t\t\u003cvmTransferDigit\u003e51\u003c/vmTransferDigit\u003e\n\t\t\u003cvmBoxType\u003eSKILL\u003c/vmBoxType\u003e\n\t\t\u003cvmSkillBox\u003e\n\t\t\t\u003cid\u003e266194\u003c/id\u003e\n\t\t\t\u003cname\u003esk_skill1\u003c/name\u003e\n\t\t\u003c/vmSkillBox\u003e\n\t\t\u003cclearDigitBuffer\u003etrue\u003c/clearDigitBuffer\u003e\n\t\t\u003cenableMusicOnHold\u003etrue\u003c/enableMusicOnHold\u003e\n\t\t\u003cannouncements\u003e\n\t\t\t\u003cenabled\u003etrue\u003c/enabled\u003e\n\t\t\t\u003cloopped\u003etrue\u003c/loopped\u003e\n\t\t\t\u003ctimeout\u003e3\u003c/timeout\u003e\n\t\t\t\u003cprompt/\u003e\n\t\t\t\u003cannType\u003eEWT\u003c/annType\u003e\n\t\t\u003c/announcements\u003e\n\t\t\u003cannouncements\u003e\n\t\t\t\u003cenabled\u003etrue\u003c/enabled\u003e\n\t\t\t\u003cloopped\u003efalse\u003c/loopped\u003e\n\t\t\t\u003ctimeout\u003e0\u003c/timeout\u003e\n\t\t\t\u003cprompt\u003e\n\t\t\t\t\u003cid\u003e222995\u003c/id\u003e\n\t\t\t\t\u003cname\u003ehold music\u003c/name\u003e\n\t\t\t\u003c/prompt\u003e\n\t\t\t\u003cannType\u003ePROMPT\u003c/annType\u003e\n\t\t\u003c/announcements\u003e\n\t\t\u003cannouncements\u003e\n\t\t\t\u003cenabled\u003efalse\u003c/enabled\u003e\n\t\t\t\u003cloopped\u003efalse\u003c/loopped\u003e\n\t\t\t\u003ctimeout\u003e1\u003c/timeout\u003e\n\t\t\t\u003cprompt/\u003e\n\t\t\t\u003cannType\u003eEWT\u003c/annType\u003e\n\t\t\u003c/announcements\u003e\n\t\t\u003cannouncements\u003e\n\t\t\t\u003cenabled\u003efalse\u003c/enabled\u003e\n\t\t\t\u003cloopped\u003efalse\u003c/loopped\u003e\n\t\t\t\u003ctimeout\u003e1\u003c/timeout\u003e\n\t\t\t\u003cprompt/\u003e\n\t\t\t\u003cannType\u003eEWT\u003c/annType\u003e\n\t\t\u003c/announcements\u003e\n\t\t\u003cannouncements\u003e\n\t\t\t\u003cenabled\u003efalse\u003c/enabled\u003e\n\t\t\t\u003cloopped\u003efalse\u003c/loopped\u003e\n\t\t\t\u003ctimeout\u003e1\u003c/timeout\u003e\n\t\t\t\u003cprompt/\u003e\n\t\t\t\u003cannType\u003eEWT\u003c/annType\u003e\n\t\t\u003c/announcements\u003e\n\t\t\u003cconnectors\u003e\n\t\t\t\u003cid\u003e119033\u003c/id\u003e\n\t\t\t\u003cname\u003eConn4_Accept\u003c/name\u003e\n\t\t\u003c/connectors\u003e\n\t\t\u003cpriorityChangeType\u003eINCREASE\u003c/priorityChangeType\u003e\n\t\t\u003cpriorityChangeValue\u003e\n\t\t\t\u003cisVarSelected\u003efalse\u003c/isVarSelected\u003e\n\t\t\t\u003cintegerValue\u003e\n\t\t\t\t\u003cvalue\u003e10\u003c/value\u003e\n\t\t\t\u003c/integerValue\u003e\n\t\t\u003c/priorityChangeValue\u003e\n\t\t\u003clistOfSkillsEx\u003e\n\t\t\t\u003cextrnalObj\u003e\n\t\t\t\t\u003cid\u003e266194\u003c/id\u003e\n\t\t\t\t\u003cname\u003esk_skill1\u003c/name\u003e\n\t\t\t\u003c/extrnalObj\u003e\n\t\t\t\u003cvarSelected\u003efalse\u003c/varSelected\u003e\n\t\t\u003c/listOfSkillsEx\u003e\n\t\t\u003clistOfSkillsEx\u003e\n\t\t\t\u003cextrnalObj\u003e\n\t\t\t\t\u003cid\u003e266254\u003c/id\u003e\n\t\t\t\t\u003cname\u003emk_skill1\u003c/name\u003e\n\t\t\t\u003c/extrnalObj\u003e\n\t\t\t\u003cvarSelected\u003efalse\u003c/varSelected\u003e\n\t\t\u003c/listOfSkillsEx\u003e\n\t\t\u003ctaskType\u003e0\u003c/taskType\u003e\n\t\t\u003ctransferAlgorithm\u003e\n\t\t\t\u003calgorithmType\u003eLongestReadyTime\u003c/algorithmType\u003e\n\t\t\t\u003cstatAlgorithmTimeWindow\u003eE15minutes\u003c/statAlgorithmTimeWindow\u003e\n\t\t\u003c/transferAlgorithm\u003e\n\t\t\u003crecordedFilesAction\u003eKEEP_AS_RECORDINGD\u003c/recordedFilesAction\u003e\n\t\t\u003ccallbackNumberToCav\u003e\n\t\t\t\u003cid\u003e201\u003c/id\u003e\n\t\t\t\u003cname\u003eqwerty\u003c/name\u003e\n\t\t\u003c/callbackNumberToCav\u003e\n\t\t\u003ccallbackNumberFromCav\u003e\n\t\t\t\u003cid\u003e63\u003c/id\u003e\n\t\t\t\u003cname\u003enumber\u003c/name\u003e\n\t\t\u003c/callbackNumberFromCav\u003e\n\t\t\u003ccallbackPhoneNumberPrompt\u003e\n\t\t\t\u003cid\u003e-26\u003c/id\u003e\n\t\t\t\u003cname\u003eQC1YourPhoneNumberIs\u003c/name\u003e\n\t\t\u003c/callbackPhoneNumberPrompt\u003e\n\t\t\u003ccallbackConfirmingPhoneNumberPrompt\u003e\n\t\t\t\u003cid\u003e-27\u003c/id\u003e\n\t\t\t\u003cname\u003eQC2PressOneIfCorrect\u003c/name\u003e\n\t\t\u003c/callbackConfirmingPhoneNumberPrompt\u003e\n\t\t\u003ccallbackEnteringPhoneNumberPrompt\u003e\n\t\t\t\u003cid\u003e-28\u003c/id\u003e\n\t\t\t\u003cname\u003eQC3PleaseEnterYourPhoneNumber\u003c/name\u003e\n\t\t\u003c/callbackEnteringPhoneNumberPrompt\u003e\n\t\t\u003ccallbackRecordingCallerNamePrompt\u003e\n\t\t\t\u003cid\u003e-36\u003c/id\u003e\n\t\t\t\u003cname\u003eQC5PleaseClearlySayYourName\u003c/name\u003e\n\t\t\u003c/callbackRecordingCallerNamePrompt\u003e\n\t\t\u003ccallbackConfirmationPrompt\u003e\n\t\t\t\u003cid\u003e-29\u003c/id\u003e\n\t\t\t\u003cname\u003eQC4YouWillBeCalled\u003c/name\u003e\n\t\t\u003c/callbackConfirmationPrompt\u003e\n\t\t\u003ccallbackQueueTimeoutSec\u003e18000\u003c/callbackQueueTimeoutSec\u003e\n\t\t\u003ccallbackEnterDigitsMaxTimeSec\u003e5\u003c/callbackEnterDigitsMaxTimeSec\u003e\n\t\t\u003ccallbackAllowInternational\u003efalse\u003c/callbackAllowInternational\u003e\n\t\t\u003cisTcpaConsentEnabled\u003etrue\u003c/isTcpaConsentEnabled\u003e\n\t\t\u003ctcpaConsentText\u003eH4sIAAAAAAAAADWOwQ3DMAwDV+EARYboL58uUPSh2EwjwJYCy878TVDkS/IOnJHcgtbRHY2JehBJ\nSgmszSvezxFqjMBLKj+Qc7gRnYX75kbYqAsb9uaHZmbI4gcfUEtlZLXv1VTv6iblL55mDMts0cXy\naTuV9wcNmHfIFWS9IPiKfbS0SXD6AQLT2xmtAAAA\u003c/tcpaConsentText\u003e\n\t\t\u003cignoreSkillsOrder\u003efalse\u003c/ignoreSkillsOrder\u003e\n\t\t\u003crecordCallerNameOnQueueCallback\u003efalse\u003c/recordCallerNameOnQueueCallback\u003e\n\t",
		},
	}
	expected.SetGeneralInfo("SkillTransfer4", "39FE8C2F1E2744D6975EF11767639A80",
		[]ivr.ModuleID{"E97CB11F6641449FB9D34E6A1C38326C"}, "E493C28EA263451CBECAA007C416F4FF", "",
		"", "false")

	exp, err1 := json.MarshalIndent(expected, "", "  ")
	stm, err2 := json.MarshalIndent(mhu, "", "  ")

	if err1 != nil || err2 != nil || string(exp) != string(stm) {
		t.Errorf("\nSkillTransfer module: \n%s \nwas expected, in reality: \n%s", string(exp), string(stm))
	}
}
