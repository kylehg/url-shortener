package shawty

import "testing"

func TestGetAndSet(t *testing.T) {
	var (
		err  error
		code string
		url  string
	)

	google := "https://www.google.com/"
	facebook := "https://www.facebook.com/"
	medium := "https://medium.com/"

	conn := getConn()
	if _, err = conn.Do("SELECT", TEST_DB); err != nil {
		panic("Cannot select test database")
	}
	conn.Do("FLUSHDB")

	if err = SetDefaultCode(google, "googl"); err != nil {
		t.Error("Setting a default code should work: " + err.Error())
	}

	if url, err = GetUrl("googl"); url != google || err != nil {
		t.Error("Getting by a default code should work: " + err.Error())
	}

	if code, err = GetDefaultCode(google); code != "googl" || err != nil {
		t.Error("Getting a default code by URL should work: " + err.Error())
	}

	if err = SetCustomCode(facebook, "f"); err != nil {
		t.Error("Setting a custom code should work: " + err.Error())
	}

	if url, err = GetUrl("f"); url != facebook || err != nil {
		t.Error("Getting by a custom code should work: " + err.Error())
	}

	if code, err = GetDefaultCode(facebook); code != "" || err == nil {
		t.Error("Getting a default code when there's only a custom code should return nil")
	}

	if err = SetDefaultCode(google, "moogl"); err == nil {
		t.Error("Setting multiple default codes should fail")
	}

	if url, err = GetUrl("moogl"); url == google || err == nil {
		t.Error("Setting multiple default codes should fail to create a code -> url mapping")
	}

	if err = SetCustomCode(facebook, "fb"); err != nil {
		t.Error("Setting multiple custom codes should work: " + err.Error())
	}

	if err = SetDefaultCode(facebook, "faceb"); err != nil {
		t.Error("Setting a default code and a custom code should work: " + err.Error())
	}

	if err = SetDefaultCode(medium, "googl"); err == nil {
		t.Error("Overwriting a default code should fail")
	}

	if url, err = GetUrl("googl"); url != google || err != nil {
		t.Error("The original default code should not be overwritten: " + err.Error())
	}

	if err = SetCustomCode(medium, "f"); err == nil {
		t.Error("Overwriting a custom code should fail")
	}

	if url, err = GetUrl("f"); url != facebook || err != nil {
		t.Error("The original custom code should not be overwritten: " + err.Error())
	}

	if err = SetDefaultCode(google, "googl"); err == nil {
		t.Error("Resetting the same default code should still fail")
	}

	if err = SetCustomCode(facebook, "fb"); err == nil {
		t.Error("Resetting the same custom code should still fail")
	}

	conn.Do("FLUSHDB")
}
