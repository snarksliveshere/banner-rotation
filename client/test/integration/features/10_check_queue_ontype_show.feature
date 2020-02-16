Feature: Test Notification After GetBanner BY GRPC method
  I send Request by GRPC to GetBanner
  I want to receive Response Status success and Banner id (string)

  Scenario: Get Notification After GetBanner
    When I send request to GRPC SendGetBannerMessage with audience "male_adult" and slot "top_slot_id"
    Then Status should be equal to success "success"
    And  The response bannerId should not be empty string
    And  Notification after SendGetBannerMessage must contain type "show" and audience "male_adult" and slot "top_slot_id"

  Scenario: Error Scenario
    When I send error request to GRPC SendGetBannerMessage with audience "fake_adult" and slot "fake_slot"
    Then Error must not be empty