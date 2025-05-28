package main

import (
	"bytes"
	"html/template"
	"net/http"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

const tmpl = `
<!DOCTYPE html>
<html lang="en">

<head>
    <meta http-equiv="Content-Type" content="text/html charset=UTF-8" />
    <meta name="format-detection" content="telephone=no,date=no,address=no,email=no,url=no" />

    <style>
        * {
            font-family: 'SF Pro Text', sans-serif;
            font-size: 18px;
            font-weight: normal;
            font-style: normal;
            font-stretch: normal;
            line-height: 1.5;
            letter-spacing: normal;
            text-align: left;
            color: #000000;
        }

        #body,
        body {
            background-color: #EAEAEA;
            margin: 0;
            padding: 0;
            width: 100%;
        }

        h1,
        h2,
        h3 {
            font-weight: bold;
            margin-top: 0;
        }

        h1 {
            font-size: 18px;
            line-height: 1.2;
            text-align: left;
            margin-bottom: 20px;
            margin-top: 40px;
        }

        h2 {
            font-size: 18px;
            line-height: 1.2;
            text-align: left;
            margin-bottom: 20px;
            margin-top: 20px;
        }

        h3 {
            font-size: 18px;
            line-height: 1.33;
        }

        h4 {
            color: #f7f4f8;
            font-size: 14px;
            line-height: 1.71;
            font-weight: normal;
            text-decoration-line: none;
        }

        p {
            text-transform: initial;
            line-height: 22px;
            margin-top: 0;
            margin-bottom: 15px;
            font-weight: 400;
            font-size: 12px;
            color: #000000A6;

        }


        b {
            font-weight: bold;
            color: #000000;

        }


        span {
            font-family: Arial;
            font-size: 18px;
            font-weight: bold;
            font-stretch: normal;
            font-style: normal;
            line-height: 1.71;
            letter-spacing: normal;
            text-align: left;
            color: #7303c0;
            text-decoration: none;
            text-underline: none;
        }

        a {
            color: #7303c0;
            font-size: 12px;
            font-family: Arial;
            font-weight: bold;
            font-stretch: normal;
            font-style: normal;
            line-height: 1.71;
            letter-spacing: normal;
            text-decoration: none;



        }




        a.gradient-button {
            background-size: cover;
            background-color: #192d59;
            color: white !important;
            padding: 17px;
            border-radius: 30px;
            text-transform: uppercase;
            text-decoration: none;
            display: inline-block;
            width: 252px;
            text-align: center;
            white-space: nowrap;
            font-size: 17px;
            font-weight: bold !important;
            line-height: 1.2;
            font-weight: 500;
            box-shadow: 1px 2px 12px #0000004D;
            -webkit-tap-highlight-color: #141C52;
            transition: all 150ms ease-in-out;
        }
    </style>
</head>

<body>
    <div id="body">
        <table style="width: 100%;border-spacing: 0;border-collapse: collapse;">

            <td style="width: 100%; padding: 0; text-align: center;">
                <table
                    style="width: 100%;max-width: 660px;border-spacing: 0;border-collapse: collapse;margin: 0 auto;background: #FFFFFF;">

                    <tr style="background: #3B2193; 
							background: linear-gradient(to right, #3B2193 0%, #AA90F1 100%);
							background-size: cover; box-shadow: 3px 3px 10px #0000004D;">

                        <td style=" width: 5%;  ">&nbsp;
                        </td>
                        <td style="height: 50px; width: 90%; text-align: left;  ">
                            <div
                                style="font-size: 1px; color: #141C52; line-height: 1px; max-height: 0px; overflow: hidden;">
                                Nuno Wallet
                            </div>
                            <img src="https://haxvpwwp.elementor.cloud/wp-content/uploads/2024/12/Vector.png" alt="" width="90" height="25">
    </td>
                        <td style="width: 5%; ">&nbsp;
                        </td>
                    </tr>



                    <td width="60px">&nbsp;</td>


                    <tr>
                        <!-- Left Section: 5% -->
                        <td width="5%">&nbsp;</td>

                        <!-- Middle Section: Smaller Content (55%) -->
                        <td width="55%" style="text-align: left; vertical-align: top; padding-right: 10px;">
                            <h1 style="margin: 0;">Dear mars</h1>
                            <p style="text-align: left; margin: 10px 0;">
                                Please find below your password reset link.
                            </p>
                            <a style="margin: 0; font-size: 18px; color: #000000; line-height: 22px; text-decoration: underline"
                                href="https://nunobusiness-frontend.vercel.app//reset-password/000000">Reset
                                Password</a>
                            <p style="text-align: left; margin: 10px 0; font-size: 18px; color: #000000;">
                                Secret code: 000000
                            </p>

                            <p style="text-align: left; margin: 10px 0;">
                                The code is valid for 10 minutes. If you did not just make a request to change your
                                password,
                                please contact us. Thank you for choosing Nuno.
                            </p>
                            <p style="text-align: left; margin: 20px 0;">
                                If you have any questions or need assistance, our customer support team is available
                                24/7.
                                You can reach us at
                                <a style="color: #00aae9; font-size: 12px;"
                                    href="mailto:support@digitalwallet.com">support@digitalwallet.com</a>
                                or call us at
                                <a style="color: #00aae9; font-size: 12px;" href="tel:+18005550199">1-800-555-0199</a>.
                            </p>
                        </td>

                        <!-- Right Section: Larger Content (40%) with Image -->
                        <td width="40%" style="text-align: right; vertical-align: top; padding-right: 10px;">
                            <img src="https://haxvpwwp.elementor.cloud/wp-content/uploads/2024/12/Email-_-Code.png"
             alt="Secure Vault"
             style="max-width: 120px; height: 120px; margin: 0;">
    </td>
                    </tr>

                    <!-- Background Image Row  -->
                    <tr>
                        <td colspan="3" style="padding: 0; margin: 0;">

                            <table role="presentation" width="100%" cellspacing="0" cellpadding="0" border="0"
                                style="margin: 0; padding: 0;">
                                <tr>
                                    <td style="padding: 0; text-align: center; line-height: 0; margin: 0;">
                                        <img src="https://haxvpwwp.elementor.cloud/wp-content/uploads/2024/12/IMG_188C9863F6D1-1-2-x2.png" alt="Background Image" width="100%" style="display: block; max-width: 660px; height: auto; margin: 0; padding: 0;">
                </td>
                                </tr>
                            </table>
                        </td>
                    </tr>


                    <!-- Footer Section -->
                    <tr>
                        <td colspan="3" style="background-color:  #8c8c8c; ">
                            <table style="width: 100%;border-spacing: 0;border-collapse: collapse;margin: 0 auto;"
                                cellspacing="0" cellpadding="0" ;>

                                <tr>
                                    <td width="5%">&nbsp;</td>
                                    <td style="width: 90%; height: 100%;" height="100%">
                                        <!-- Empty Spacer -->
                                        <p style="margin: 0; line-height: 20px; font-size: 20px;">&nbsp;</p>

                                        <!-- Main Footer Table -->
                                        <table style="width: 100%; border-spacing: 0; border-collapse: collapse;">
                                            <tr>
                                                <!-- Left Content -->
                                                <td style="text-align: left; vertical-align: middle; padding: 0;">
                                                    <p
                                                        style="margin: 0; font-size: 10px; color: #ffffff; line-height: 1.5;">
                                                        For more information visit
                                                        <a style="margin: 0; font-size: 10px; color: #ffffff; line-height: 22px; text-decoration: underline"
                                                            href="https://nucleus-wallet-app.vercel.app/">nuno.com</a>
                                                    </p>
                                                </td>

                                                <!-- Right Content (Icons) -->
                                                <td style="text-align: right; vertical-align: middle; padding: 0;">
                                                    <table
                                                        style="border-spacing: 0; border-collapse: collapse; display: inline-block;">
                                                        <tr>
                                                            <td style="padding: 0 10px;">
                                                                <a href="#" style="display: inline-block;">
                                                                    <img src="https://haxvpwwp.elementor.cloud/wp-content/uploads/2024/12/Vector-x2-yt.png"
                         alt="YouTube" width="14" height="10" style="display: inline-block; border: none;">
                </a>
                                                            </td>
                                                            <td style="padding: 0 10px;">
                                                                <a href="#" style="display: inline-block;">
                                                                    <img src="https://haxvpwwp.elementor.cloud/wp-content/uploads/2024/12/Vector-in-x2.png"
                         alt="LinkedIn" width="13" height="13" style="display: inline-block; border: none;">
                </a>
                                                            </td>
                                                            <td style="padding: 0 10px;">
                                                                <a href="#" style="display: inline-block;">
                                                                    <img src="https://haxvpwwp.elementor.cloud/wp-content/uploads/2024/12/Vector-x-x2.png"
                         alt="X" width="13" height="13" style="display: inline-block; border: none;">
                </a>
                                                            </td>
                                                        </tr>
                                                    </table>
                                                </td>

                                            </tr>
                                        </table>

                                        <!-- Spacer -->
                                        <p style="margin: 0; line-height: 20px;">&nbsp;</p>

                                        <p style="margin: 0; font-size: 10px; color: #ffffff; line-height: 22px;">
                                            Terms and Conditions apply. We NEVER ask our merchants or our users to send
                                            the Personal <br>
					Account Number (PAN) for Credit / Debit cards, the CVV security code, online password, online PIN <br>
					or mobile PIN via email or when identifying you via the phone.
                </p>
                                        <!-- Spacer -->
                                        <p style="margin: 0; line-height: 20px;">&nbsp;</p>

                                    </td>



                                    <td rowspan="1" style="width: 5%; height:100px; " width="5%" height="100%">
                                        <table style="width: 100%; height:100%;" width="100%" cellspacing="0"
                                            cellpadding="0" height="100%">
                                            <tr>

                                                <td style=" text-align: right; width: 20px; height: 100%" width="20px"
                                                    height="100%">
                                                    <a style="border: 0; text-decoration:none;">

                                                    </a>
                                                </td>
                                            </tr>
                                        </table>
                                    </td>
                        </td>

            </td>
        </table>
        </tr>
        </td>
        </table>
        </tr>

        </table>
        </table>
    </div>
</body>

</html>
`

type PageData struct {
	Name string
}

func handler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.New("hello").Parse(tmpl))
	data := PageData{Name: "World"}
	// Render HTML to a buffer
	var buf bytes.Buffer
	t.Execute(&buf, data)

	// Use wkhtmltopdf to convert HTML to PDF
	// Requires github.com/SebastiaanKlippert/go-wkhtmltopdf and wkhtmltopdf installed on your system

	pdf, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		http.Error(w, "Failed to initialize PDF generator", http.StatusInternalServerError)
		return
	}

	page := wkhtmltopdf.NewPageReader(bytes.NewReader(buf.Bytes()))
	pdf.AddPage(page)

	err = pdf.Create()
	if err != nil {
		http.Error(w, "Failed to generate PDF", http.StatusInternalServerError)
		return
	}

	// Set headers for PDF download
	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", "attachment; filename=\"output.pdf\"")
	w.Write(pdf.Bytes())
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
