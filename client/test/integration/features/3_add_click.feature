Feature: AddClick By GRPC
  I send Request by GRPC to AddClick
  And I send banner, slot & audience params
  I want to receive Response Status success

  Scenario: AddClick
    When I send request to GRPC SendAddClickBannerMessage with banner "some_male2_adult_app_id" and slot "top_slot_id" and audience "male_adult"
    Then Status should be equal to success "success"

  Scenario: Error Scenario
    When I send error request to GRPC SendAddClickBannerMessage with banner "fake_banner" and slot "fake_slot" and audience "fake_audience"
    Then Error must not be empty