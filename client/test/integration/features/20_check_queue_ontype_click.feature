Feature: Test Notification After AddClick BY GRPC method
  I send Request by GRPC to GetBanner
  I want to receive Response Status success and Banner id (string)

  Scenario: Get Notification After AddClick
    When I send request to GRPC SendAddClickBannerMessage with banner "some_male2_adult_app_id" and slot "top_slot_id" and audience "male_adult"
    Then Status should be equal to success "success"
    And  The response bannerId should not be empty string
    And  Notification SendAddClickBannerMessage must contain type "click" and banner "some_male2_adult_app_id" and slot "top_slot_id" and audience "male_adult"

  Scenario: Error Scenario
    When I send error request to GRPC SendAddClickBannerMessage with banner "fake_banner" and slot "fake_slot" and audience "fake_audience"
    Then Error must not be empty
