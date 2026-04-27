package auth

import (
	"fmt"
	"html"
	"strings"
	"time"
)

type authEmailTemplateData struct {
	Preview      string
	StatusKey    string
	StatusValue  string
	ActionKey    string
	ActionValue  string
	Title        string
	Subtitle     string
	Name         string
	Body         string
	ButtonLabel  string
	ButtonURL    string
	ExpiryText   string
	FooterNote   string
	AccentColor  string
	AccentShadow string
	ButtonColor  string
	ButtonText   string
}

func passwordResetEmailHTML(name string, resetURL string, ttl time.Duration) string {
	return renderAuthEmailTemplate(authEmailTemplateData{
		Preview:      "Reset your PIXEL.CMS admin password.",
		StatusKey:    "recovery_mode",
		StatusValue:  "true",
		ActionKey:    "reset_token",
		ActionValue:  "queued_for_dispatch",
		Title:        "PASSWORD RESET",
		Subtitle:     "~ ADMIN TERMINAL v4.20",
		Name:         name,
		Body:         "A password reset was requested for your account. Use the secure access link below to set a new password.",
		ButtonLabel:  "RESET PASSWORD",
		ButtonURL:    resetURL,
		ExpiryText:   fmt.Sprintf("This reset link expires in %d minutes.", ttlMinutes(ttl)),
		FooterNote:   "If you did not request this, ignore this email and your password will stay unchanged.",
		AccentColor:  "#f1fa8c",
		AccentShadow: "rgba(241,250,140,0.45)",
		ButtonColor:  "#50fa7b",
		ButtonText:   "#191a27",
	})
}

func verificationEmailHTML(name string, verificationURL string, ttl time.Duration) string {
	return renderAuthEmailTemplate(authEmailTemplateData{
		Preview:      "Verify your PIXEL.CMS admin email.",
		StatusKey:    "email_verification",
		StatusValue:  "true",
		ActionKey:    "verify_link",
		ActionValue:  "ready_to_open",
		Title:        "VERIFY EMAIL",
		Subtitle:     "~ ADMIN TERMINAL v4.20",
		Name:         name,
		Body:         "Please verify your email address for your admin account. Use the secure verification link below to continue.",
		ButtonLabel:  "VERIFY EMAIL",
		ButtonURL:    verificationURL,
		ExpiryText:   fmt.Sprintf("This verification link expires in %d minutes.", ttlMinutes(ttl)),
		FooterNote:   "If you did not request this, ignore this email and no account changes will be made.",
		AccentColor:  "#50fa7b",
		AccentShadow: "rgba(80,250,123,0.45)",
		ButtonColor:  "#8be9fd",
		ButtonText:   "#191a27",
	})
}

func renderAuthEmailTemplate(data authEmailTemplateData) string {
	name := strings.TrimSpace(data.Name)
	if name == "" {
		name = "Admin"
	}

	return fmt.Sprintf(`<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width,initial-scale=1">
  <title>%s</title>
</head>
<body style="margin:0;padding:0;background:#282a36;color:#f8f8f2;font-family:'Courier New',Courier,monospace;">
  <div style="display:none;max-height:0;overflow:hidden;opacity:0;color:transparent;">%s</div>
  <table role="presentation" width="100%%" cellspacing="0" cellpadding="0" style="width:100%%;border-collapse:collapse;background:#282a36;background-image:linear-gradient(rgba(255,255,255,0.025) 1px,transparent 1px),linear-gradient(90deg,rgba(255,255,255,0.025) 1px,transparent 1px);background-size:24px 24px;">
    <tr>
      <td align="center" style="padding:32px 16px;">
        <table role="presentation" width="100%%" cellspacing="0" cellpadding="0" style="width:100%%;max-width:520px;border-collapse:collapse;background:#1f2030;border:4px solid #44475a;box-shadow:0 0 0 4px #191a27,12px 12px 0 #12131c;">
          <tr>
            <td style="padding:26px 26px 16px;">
              <table role="presentation" cellspacing="0" cellpadding="0" style="border-collapse:collapse;">
                <tr>
                  <td width="72" valign="top" style="width:72px;padding-right:16px;">
                    <table role="presentation" width="72" height="72" cellspacing="0" cellpadding="0" style="width:72px;height:72px;border-collapse:collapse;background:#bd93f9;border-top:6px solid #d4b3ff;border-left:6px solid #d4b3ff;border-right:6px solid #7a5cb8;border-bottom:6px solid #7a5cb8;">
                      <tr><td style="height:18px;background:#bdc0cc;border-right:18px solid #44475a;"></td></tr>
                      <tr><td style="height:22px;"></td></tr>
                      <tr><td style="height:20px;background:#f8f8f2;border-left:14px solid #bd93f9;border-right:14px solid #bd93f9;"></td></tr>
                    </table>
                  </td>
                  <td valign="middle">
                    <div style="margin:0 0 9px;color:#bd93f9;font-family:'Courier New',Courier,monospace;font-size:26px;line-height:1.05;font-weight:800;letter-spacing:0;text-shadow:3px 3px 0 #7a5cb8;">PIXEL.CMS</div>
                    <div style="margin:0;color:#8be9fd;font-size:18px;line-height:1.2;">%s</div>
                  </td>
                </tr>
              </table>
            </td>
          </tr>
          <tr>
            <td style="padding:0 26px 18px;">
              <div style="height:8px;background:linear-gradient(90deg,#bd93f9 0 25%%,transparent 25%% 37.5%%,#8be9fd 37.5%% 62.5%%,transparent 62.5%% 75%%,#ff79c6 75%% 100%%);background-size:64px 8px;"></div>
            </td>
          </tr>
          <tr>
            <td style="padding:0 26px 18px;color:#8be9fd;font-size:19px;line-height:1.32;">
              &gt; <span style="color:#f8f8f2;">%s</span>: %s<br>
              &gt; <span style="color:#f8f8f2;">%s</span>: %s <span style="display:inline-block;width:9px;height:17px;background:#50fa7b;vertical-align:text-bottom;"></span>
            </td>
          </tr>
          <tr>
            <td style="padding:0 26px 22px;">
              <div style="padding:11px 12px;color:%s;background:rgba(255,255,255,0.035);box-shadow:inset 0 0 0 2px %s;font-size:12px;line-height:1.65;font-weight:700;letter-spacing:0;">%s</div>
            </td>
          </tr>
          <tr>
            <td style="padding:0 26px 8px;">
              <h1 style="margin:0 0 14px;color:#f8f8f2;font-size:24px;line-height:1.25;font-weight:800;letter-spacing:0;">%s</h1>
              <p style="margin:0 0 12px;color:#f8f8f2;font-size:16px;line-height:1.65;">Hello <strong style="color:#8be9fd;">%s</strong>,</p>
              <p style="margin:0;color:#f8f8f2;font-size:16px;line-height:1.65;">%s</p>
            </td>
          </tr>
          <tr>
            <td align="center" style="padding:22px 26px 20px;">
              <a href="%s" style="display:block;padding:16px 18px;background:%s;color:%s;border:3px solid #44475a;box-shadow:5px 5px 0 #12131c;text-decoration:none;font-size:13px;line-height:1.3;font-weight:800;letter-spacing:0;">%s</a>
            </td>
          </tr>
          <tr>
            <td style="padding:0 26px 22px;">
              <p style="margin:0 0 10px;color:#f1fa8c;font-size:13px;line-height:1.6;">%s</p>
              <p style="margin:0;color:#777a95;font-size:12px;line-height:1.7;">%s</p>
            </td>
          </tr>
          <tr>
            <td align="center" style="padding:20px 26px 24px;border-top:3px solid #44475a;">
              <p style="margin:0;color:#777a95;font-size:11px;line-height:1.6;letter-spacing:1px;">PORTFOLIO ADMIN SYSTEM - MADE WITH &lt;3 &amp; PIXELS</p>
            </td>
          </tr>
        </table>
      </td>
    </tr>
  </table>
</body>
</html>`,
		html.EscapeString(data.Title),
		html.EscapeString(data.Preview),
		html.EscapeString(data.Subtitle),
		html.EscapeString(data.StatusKey),
		html.EscapeString(data.StatusValue),
		html.EscapeString(data.ActionKey),
		html.EscapeString(data.ActionValue),
		data.AccentColor,
		data.AccentShadow,
		html.EscapeString(data.Title),
		html.EscapeString(data.Title),
		html.EscapeString(name),
		html.EscapeString(data.Body),
		html.EscapeString(data.ButtonURL),
		data.ButtonColor,
		data.ButtonText,
		html.EscapeString(data.ButtonLabel),
		html.EscapeString(data.ExpiryText),
		html.EscapeString(data.FooterNote),
	)
}

func ttlMinutes(ttl time.Duration) int {
	minutes := int(ttl.Minutes())
	if minutes < 1 {
		return 1
	}

	return minutes
}
