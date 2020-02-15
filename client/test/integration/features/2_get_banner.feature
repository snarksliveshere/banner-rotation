Feature: GetBanner BY GRPC method
  I send Request by GRPC to GetBanner
  I want to receive Response Status success and Banner id (string)

  Scenario: Get Banner
    When I send request to GRPC server SendGetBannerMessage
    Then Status should be equal to success "success"
    And  The response bannerId should not be empty string

