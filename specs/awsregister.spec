# Model
model:
  rest_name: awsregister
  resource_name: awsregister
  entity_name: AWSRegister
  package: bill
  group: none
  description: This api allows AWS customer to register with Aporeto SaaS for billing.
  extends:
  - '@identifiable-stored'
  - '@timeable'

# Attributes
attributes:
  v1:
  - name: provider
    description: Token Provided by AWS.
    type: string
    exposed: true
