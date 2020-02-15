Feature: Check for stability GRPC server
  I send check request by GRPC
  I want to receive Response Status success

  Scenario: HealthCheck
    When I send request to GRPC SendHealthCheckMessage
    Then Status should be equal to success "success"

