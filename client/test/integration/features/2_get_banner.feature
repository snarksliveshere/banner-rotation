Feature: GetBanner BY GRPC method
  I send Request by GRPC to GetBanner
  I want to receive Response Status success and Banner id (string)

  Scenario: Get Banner
    When I send request to GRPC SendGetBannerMessage with audience "male_adult" and slot "top_slot_id"
    Then Status should be equal to success "success"
    And  The response bannerId should not be empty string

# other way