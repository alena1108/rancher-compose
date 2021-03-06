package client

const (
	MACHINE_TYPE = "machine"
)

type Machine struct {
	Resource

	AccountId string `json:"accountId,omitempty" yaml:"account_id,omitempty"`

	Amazonec2Config *Amazonec2Config `json:"amazonec2Config,omitempty" yaml:"amazonec2config,omitempty"`

	AuthCertificateAuthority string `json:"authCertificateAuthority,omitempty" yaml:"auth_certificate_authority,omitempty"`

	AuthKey string `json:"authKey,omitempty" yaml:"auth_key,omitempty"`

	AzureConfig *AzureConfig `json:"azureConfig,omitempty" yaml:"azure_config,omitempty"`

	Created string `json:"created,omitempty" yaml:"created,omitempty"`

	Data map[string]interface{} `json:"data,omitempty" yaml:"data,omitempty"`

	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	DigitaloceanConfig *DigitaloceanConfig `json:"digitaloceanConfig,omitempty" yaml:"digitalocean_config,omitempty"`

	Driver string `json:"driver,omitempty" yaml:"driver,omitempty"`

	ExoscaleConfig *ExoscaleConfig `json:"exoscaleConfig,omitempty" yaml:"exoscale_config,omitempty"`

	ExternalId string `json:"externalId,omitempty" yaml:"external_id,omitempty"`

	Kind string `json:"kind,omitempty" yaml:"kind,omitempty"`

	Labels map[string]interface{} `json:"labels,omitempty" yaml:"labels,omitempty"`

	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	OpenstackConfig *OpenstackConfig `json:"openstackConfig,omitempty" yaml:"openstack_config,omitempty"`

	PacketConfig *PacketConfig `json:"packetConfig,omitempty" yaml:"packet_config,omitempty"`

	RackspaceConfig *RackspaceConfig `json:"rackspaceConfig,omitempty" yaml:"rackspace_config,omitempty"`

	RemoveTime string `json:"removeTime,omitempty" yaml:"remove_time,omitempty"`

	Removed string `json:"removed,omitempty" yaml:"removed,omitempty"`

	SoftlayerConfig *SoftlayerConfig `json:"softlayerConfig,omitempty" yaml:"softlayer_config,omitempty"`

	State string `json:"state,omitempty" yaml:"state,omitempty"`

	Transitioning string `json:"transitioning,omitempty" yaml:"transitioning,omitempty"`

	TransitioningMessage string `json:"transitioningMessage,omitempty" yaml:"transitioning_message,omitempty"`

	TransitioningProgress int64 `json:"transitioningProgress,omitempty" yaml:"transitioning_progress,omitempty"`

	Uuid string `json:"uuid,omitempty" yaml:"uuid,omitempty"`

	VirtualboxConfig *VirtualboxConfig `json:"virtualboxConfig,omitempty" yaml:"virtualbox_config,omitempty"`

	VmwarevcloudairConfig *VmwarevcloudairConfig `json:"vmwarevcloudairConfig,omitempty" yaml:"vmwarevcloudair_config,omitempty"`

	VmwarevsphereConfig *VmwarevsphereConfig `json:"vmwarevsphereConfig,omitempty" yaml:"vmwarevsphere_config,omitempty"`
}

type MachineCollection struct {
	Collection
	Data []Machine `json:"data,omitempty"`
}

type MachineClient struct {
	rancherClient *RancherClient
}

type MachineOperations interface {
	List(opts *ListOpts) (*MachineCollection, error)
	Create(opts *Machine) (*Machine, error)
	Update(existing *Machine, updates interface{}) (*Machine, error)
	ById(id string) (*Machine, error)
	Delete(container *Machine) error

	ActionBootstrap(*Machine) (*PhysicalHost, error)

	ActionCreate(*Machine) (*PhysicalHost, error)

	ActionError(*Machine) (*PhysicalHost, error)

	ActionRemove(*Machine) (*PhysicalHost, error)

	ActionUpdate(*Machine) (*PhysicalHost, error)
}

func newMachineClient(rancherClient *RancherClient) *MachineClient {
	return &MachineClient{
		rancherClient: rancherClient,
	}
}

func (c *MachineClient) Create(container *Machine) (*Machine, error) {
	resp := &Machine{}
	err := c.rancherClient.doCreate(MACHINE_TYPE, container, resp)
	return resp, err
}

func (c *MachineClient) Update(existing *Machine, updates interface{}) (*Machine, error) {
	resp := &Machine{}
	err := c.rancherClient.doUpdate(MACHINE_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *MachineClient) List(opts *ListOpts) (*MachineCollection, error) {
	resp := &MachineCollection{}
	err := c.rancherClient.doList(MACHINE_TYPE, opts, resp)
	return resp, err
}

func (c *MachineClient) ById(id string) (*Machine, error) {
	resp := &Machine{}
	err := c.rancherClient.doById(MACHINE_TYPE, id, resp)
	return resp, err
}

func (c *MachineClient) Delete(container *Machine) error {
	return c.rancherClient.doResourceDelete(MACHINE_TYPE, &container.Resource)
}

func (c *MachineClient) ActionBootstrap(resource *Machine) (*PhysicalHost, error) {

	resp := &PhysicalHost{}

	err := c.rancherClient.doAction(MACHINE_TYPE, "bootstrap", &resource.Resource, nil, resp)

	return resp, err
}

func (c *MachineClient) ActionCreate(resource *Machine) (*PhysicalHost, error) {

	resp := &PhysicalHost{}

	err := c.rancherClient.doAction(MACHINE_TYPE, "create", &resource.Resource, nil, resp)

	return resp, err
}

func (c *MachineClient) ActionError(resource *Machine) (*PhysicalHost, error) {

	resp := &PhysicalHost{}

	err := c.rancherClient.doAction(MACHINE_TYPE, "error", &resource.Resource, nil, resp)

	return resp, err
}

func (c *MachineClient) ActionRemove(resource *Machine) (*PhysicalHost, error) {

	resp := &PhysicalHost{}

	err := c.rancherClient.doAction(MACHINE_TYPE, "remove", &resource.Resource, nil, resp)

	return resp, err
}

func (c *MachineClient) ActionUpdate(resource *Machine) (*PhysicalHost, error) {

	resp := &PhysicalHost{}

	err := c.rancherClient.doAction(MACHINE_TYPE, "update", &resource.Resource, nil, resp)

	return resp, err
}
