Feature: DeleteBannerFromSlot By GRPC
  I send Request by GRPC to sendDeleteBannerFromSlotMessage
  And I send banner & slot param
  I want to receive Response Status success

  Scenario: Success Scenario
    When I send request to GRPC sendDeleteBannerFromSlotMessage with banner "some_male2_kid_app_id" and slot "top_slot_id"
    Then Status should be equal to success "success"

  Scenario: Error Scenario
    When I send error request to GRPC sendDeleteBannerFromSlotMessage with banner "fake_banner" and slot "fake_slot"
    Then Error must not be empty