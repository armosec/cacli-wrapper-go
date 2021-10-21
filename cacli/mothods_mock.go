package cacli

import (
	"github.com/armosec/armoapi-go/armotypes"
	"github.com/armosec/opa-utils/reporthandling"

	"github.com/armosec/utils-k8s-go/secrethandling"
)

// Cacli commands
type CacliMock struct {
	backendURL  string
	credentials CredStruct
}

// NewCacli -
func NewCacliMock(backendURL string) *CacliMock {
	// Load credentials from mounted secret
	return &CacliMock{
		backendURL:  backendURL,
		credentials: CredStruct{},
	}
}

// ================================================================================================
// ================================ BASIC =========================================================
// ================================================================================================

// Login cacli login
func (cacli *CacliMock) Login() error {
	return nil
}

// Status cacli --status
func (caclim *CacliMock) Status() (*Status, error) {
	return &Status{}, nil
}

// Sign command
func (caclim *CacliMock) Sign(wlid, user, password, ociImageURL string) error {
	return nil
}

// ================================================================================================
// ================================== WT ==========================================================
// ================================================================================================

// Create command
func (caclim *CacliMock) WTCreate(wt *WorkloadTemplate, fileName string) (string, error) {
	return "", nil
}

// Apply command
func (caclim *CacliMock) WTApply(wt *WorkloadTemplate, fileName string) (string, error) {
	return "", nil

}

// Update command
func (caclim *CacliMock) WTUpdate(wt *WorkloadTemplate, fileName string) (string, error) {
	return "", nil

}

// Triplet command
func (caclim *CacliMock) WTTriplet(wlid string) (*GUIDTriplet, error) {
	return &GUIDTriplet{}, nil
}

// Get command
// func (caclim *CacliMock) Get(wlid string) error {
func (caclim *CacliMock) WTGet(wlid string) (*WorkloadTemplate, error) {
	return &WorkloadTemplate{}, nil
}

// Get command
func (caclim *CacliMock) WTDelete(wlid string) error {
	return nil
}

// Download command
func (caclim *CacliMock) WTDownload(wlid, containerName, output string) error {
	return nil
}

// Sign command
func (caclim *CacliMock) WTSign(wlid, user, password, ociImageURL string) error {
	return nil
}

// ================================================================================================
// ================================= K8S ==========================================================
// ================================================================================================

// AttachNameSpace command attach all workloads in namespace
func (caclim *CacliMock) K8SAttach(cluster, ns, wlid string, _ bool) error {
	return nil
}

// ================================================================================================
// ===================== K8S POSTURE EXCEPTION ====================================================
// ================================================================================================
func (caclim *CacliMock) K8SPEList() ([]string, error) {
	return []string{}, nil

}

func (caclim *CacliMock) K8SPEGet(name, framework, control, cluster, namespace string) ([]armotypes.PostureExceptionPolicy, error) {
	pe := []armotypes.PostureExceptionPolicy{}
	return pe, nil
}

// ================================================================================================
// ============================= K8S POSTURE ======================================================
// ================================================================================================
func (caclim *CacliMock) K8SPOSTURECreate(framework, cluster, namespace, wlid string) error {
	return nil
}

// ================================================================================================
// ============================ OPA FRAMEWORK =====================================================
// ================================================================================================

// OPAFRAMEWORKGet cacli opa get
func (caclim *CacliMock) OPAFRAMEWORKGet(name string) ([]reporthandling.Framework, error) {
	return []reporthandling.Framework{}, nil
}

// OPAFRAMEWORKDelete cacli opa delete
func (caclim *CacliMock) OPAFRAMEWORKDelete(name string) error {
	return nil
}

// OPAFRAMEWORKList - cacli opa list
func (caclim *CacliMock) OPAFRAMEWORKList() ([]string, error) {
	return []string{}, nil
}

// OPAFRAMEWORKCreate - cacli opa create
func (caclim *CacliMock) OPAFRAMEWORKCreate(framework *reporthandling.Framework, fileName string) (*reporthandling.Framework, error) {
	return nil, nil
}

// OPAFRAMEWORKUpdate - cacli opa update
func (caclim *CacliMock) OPAFRAMEWORKUpdate(framework *reporthandling.Framework, fileName string) (*reporthandling.Framework, error) {
	return nil, nil
}

// ================================================================================================
// ============================ OPA CONTROL =======================================================
// ================================================================================================

// OPAFRAMEWORKGet cacli opa get
func (caclim *CacliMock) OPACONTROLGet(name string) ([]reporthandling.Control, error) {
	return []reporthandling.Control{}, nil
}

// OPAFRAMEWORKGet cacli opa get
func (caclim *CacliMock) OPACONTROLDelete(name string) error {
	return nil
}

// OPAFRAMEWORKList - cacli opa list
func (caclim *CacliMock) OPACONTROLList() ([]string, error) {
	return []string{}, nil
}

// OPAFRAMEWORKCreate - cacli opa create
func (caclim *CacliMock) OPACONTROLCreate(control *reporthandling.Control, fileName string) (*reporthandling.Control, error) {
	return nil, nil
}

// OPAFRAMEWORKUpdate - cacli opa update
func (caclim *CacliMock) OPACONTROLUpdate(control *reporthandling.Control, fileName string) (*reporthandling.Control, error) {
	return nil, nil
}

// ================================================================================================
// ============================== OPA RULE ========================================================
// ================================================================================================

// OPAFRAMEWORKGet cacli opa get
func (caclim *CacliMock) OPARULEGet(name string) ([]reporthandling.PolicyRule, error) {
	return []reporthandling.PolicyRule{}, nil
}

// OPAFRAMEWORKGet cacli opa get
func (caclim *CacliMock) OPARULEDelete(name string) error {
	return nil
}

// OPAFRAMEWORKList - cacli opa list
func (caclim *CacliMock) OPARULEList() ([]string, error) {
	return []string{}, nil
}

// OPAFRAMEWORKCreate - cacli opa create
func (caclim *CacliMock) OPARULECreate(rule *reporthandling.PolicyRule, fileName string) (*reporthandling.PolicyRule, error) {
	return nil, nil
}

// OPAFRAMEWORKUpdate - cacli opa update
func (caclim *CacliMock) OPARULEUpdate(rule *reporthandling.PolicyRule, fileName string) (*reporthandling.PolicyRule, error) {
	return nil, nil
}

// ================================================================================================
// ================================ SECP ==========================================================
// ================================================================================================

// SecretEncrypt -
func (caclim *CacliMock) SECPEncrypt(message, inputFile, outputFile, keyID string, base64Enc bool) ([]byte, error) {
	return []byte{}, nil
}

// SecretDecrypt -
func (caclim *CacliMock) SECPDecrypt(message, inputFile, outputFile string, base64Enc bool) ([]byte, error) {
	return []byte{}, nil

}

// GetSecretAccessPolicy -
func (caclim *CacliMock) SECPGet(sid, name, cluster, namespace string) ([]secrethandling.SecretAccessPolicy, error) {
	return []secrethandling.SecretAccessPolicy{}, nil
}

// ================================================================================================
// ================================ UTILS =========================================================
// ================================================================================================

func (caclim *CacliMock) UTILSCleanup(wlid string, _ bool) error {
	return nil
}
