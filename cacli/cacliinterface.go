package cacli

import (
	"github.com/armosec/armoapi-go/armotypes"
	"github.com/armosec/opa-utils/reporthandling"
	"github.com/armosec/utils-k8s-go/secrethandling"
)

/*
Please follow the convention:
cacli <group1> <group2> <command>
The function name should look like:
GROUP1GROUP2Command (groups should be upper-case. command - first leeter upper case)

Examples:
cacli wt get -> WTGet
cacli wt Triplet -> WTTriplet
cacli secp list -> SECPList
cacli secp encrypt -> SECPEncrypt
cacli k8s attach -> K8SAttach
cacli opa framework get -> OPAFRAMEWORKGet
*/
type ICacli interface {
	// basic commands
	Login() error
	Status() (*Status, error)
	Sign(wlid, user, password, ociImageURL string) error

	// wt
	WTCreate(*WorkloadTemplate, string) (string, error)
	WTApply(*WorkloadTemplate, string) (string, error)
	WTUpdate(*WorkloadTemplate, string) (string, error)
	WTDelete(string) error
	WTTriplet(string) (*GUIDTriplet, error)
	WTGet(string) (*WorkloadTemplate, error)
	WTDownload(wlid, containerName, output string) error
	WTSign(wlid, user, password, ociImageURL string) error

	// sp
	// SPGet(name string) (*SigningProfile, error)
	// SPCreate(sp *SigningProfile) (string, error)
	// SPDelete(name string) error
	// SPGenarate(name string) (*SigningProfile, error)

	// k8s
	K8SAttach(_, _, _ string, injectLabel bool) error

	// posture
	K8SPOSTURECreate(_, _, _, _ string) error

	// posture-exceptions
	K8SPEList() ([]string, error)
	K8SPEGet(name, framework, control, cluster, namespace string) ([]armotypes.PostureExceptionPolicy, error)

	//customer config
	ClusterConfigGet(name, cluster, scope string) (armotypes.CustomerConfig, error)

	// OPA FRAMEWORK
	OPAFRAMEWORKCreate(*reporthandling.Framework, string) (*reporthandling.Framework, error)
	OPAFRAMEWORKUpdate(*reporthandling.Framework, string) (*reporthandling.Framework, error)
	OPAFRAMEWORKGet(string, bool) ([]reporthandling.Framework, error)
	OPAFRAMEWORKList(bool) ([]string, error)
	OPAFRAMEWORKDelete(string) error

	// OPA CONTROL
	OPACONTROLCreate(*reporthandling.Control, string) (*reporthandling.Control, error)
	OPACONTROLUpdate(*reporthandling.Control, string) (*reporthandling.Control, error)
	OPACONTROLGet(string) ([]reporthandling.Control, error)
	OPACONTROLList() ([]string, error)
	OPACONTROLDelete(string) error

	// OPA RULE
	OPARULECreate(*reporthandling.PolicyRule, string) (*reporthandling.PolicyRule, error)
	OPARULEUpdate(*reporthandling.PolicyRule, string) (*reporthandling.PolicyRule, error)
	OPARULEGet(string) ([]reporthandling.PolicyRule, error)
	OPARULEList() ([]string, error)
	OPARULEDelete(string) error

	// // key
	// KEYGet(string) (*Key, error)

	// secret policy

	SECPGet(sid, name, cluster, namespace string) ([]secrethandling.SecretAccessPolicy, error)
	SECPEncrypt(message, inputFile, outputFile, keyID string, base64Enc bool) ([]byte, error)
	SECPDecrypt(message, inputFile, outputFile string, base64Enc bool) ([]byte, error)
	// SECPMetadata(string, bool) (*SecretMetadata, error)
	// SECPCreate(*secrethandling.SecretAccessPolicy) (*secrethandling.SecretAccessPolicy, error)
	// SECPUpdate(*secrethandling.SecretAccessPolicy) (*secrethandling.SecretAccessPolicy, error)
	// SECPList() ([]string, error)

	// Utils
	UTILSCleanup(string, bool) error
}
