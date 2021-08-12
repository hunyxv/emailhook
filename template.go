package emailhook

import (
	"html/template"
	"io"
	"time"

	"github.com/hunyxv/utils/emailnotify"
)

var _ emailnotify.MessageImplementer = (*DefaultTemplate)(nil)

// DefaultTemplate 默认模板
type DefaultTemplate struct {
	Title   string                 // 标题
	T       string                 // 时间
	Detail  map[string]interface{} // 日志携带的额外数据
	Message string                 // 日志 msg
	Stack   map[string]string      // 堆栈

	tempalte string // 模板
}

// NewDefaultTemplate .
func NewDefaultTemplate(title, message string, t time.Time, stack map[string]string, detail map[string]interface{}) *DefaultTemplate {

	return &DefaultTemplate{
		Title:    title,
		Detail:   detail,
		Message:  message,
		Stack:    stack,
		T:        t.Format("2006-01-02 15:04:05"),
		tempalte: html,
	}
}

// Content MessageImplementer
func (t *DefaultTemplate) Content(wr io.Writer) error {
	tmpl, err := template.New("mail").Parse(t.tempalte)
	if err != nil {
		panic(err)
	}
	return tmpl.Execute(wr, t)
}

var html = `<!DOCTYPE html
PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">

<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
<!-- <title>A Simple Responsive HTML Email</title> -->
<style type="text/css">
	body {
		margin: 0;
		padding: 0;
		min-width: 100% !important;
	}

	img {
		height: auto;
	}

	.content {
		width: 100%;
		max-width: 600px;
	}

	.header {
		padding: 40px 30px 20px 30px;
	}

	.innerpadding {
		padding: 30px 30px 30px 30px;
	}

	.borderbottom {
		border-bottom: 1px solid #f2eeed;
	}

	.subhead {
		font-size: 15px;
		color: #f00808;
		font-family: sans-serif;
		/* letter-spacing: 10px; */
	}

	.h1,
	.h2,
	.bodycopy {
		color: #153643;
		font-family: sans-serif;
	}

	.h1 {
		font-size: 33px;
		line-height: 38px;
		font-weight: bold;
	}

	.h2 {
		padding: 0 0 15px 0;
		font-size: 24px;
		line-height: 28px;
		font-weight: bold;
	}

	.bodycopy {
		font-size: 16px;
		line-height: 22px;
	}

	.button {
		text-align: center;
		font-size: 18px;
		font-family: sans-serif;
		font-weight: bold;
		padding: 0 30px 0 30px;
	}

	.button a {
		color: #ffffff;
		text-decoration: none;
	}

	.footer {
		padding: 20px 30px 15px 30px;
	}

	.footercopy {
		font-family: sans-serif;
		font-size: 14px;
		color: #ffffff;
	}

	.footercopy a {
		color: #ffffff;
		text-decoration: underline;
	}

	@media only screen and (max-width: 550px),
	screen and (max-device-width: 550px) {
		body[yahoo] .hide {
			display: none !important;
		}

		body[yahoo] .buttonwrapper {
			background-color: transparent !important;
		}

		body[yahoo] .button {
			padding: 0px !important;
		}

		body[yahoo] .button a {
			background-color: #e05443;
			padding: 15px 15px 13px !important;
		}

		body[yahoo] .unsubscribe {
			display: block;
			margin-top: 20px;
			padding: 10px 50px;
			background: #2f3942;
			border-radius: 5px;
			text-decoration: none !important;
			font-weight: bold;
		}
	}

	/*@media only screen and (min-device-width: 601px) {
.content {width: 600px !important;}
.col425 {width: 425px!important;}
.col380 {width: 380px!important;}
}*/
</style>
</head>

<body yahoo bgcolor="#f6f8f1" style="margin: 0; padding: 0; min-width: 100% !important;">
<table width="100%" bgcolor="#f6f8f1" border="0" cellpadding="0" cellspacing="0">
	<tr>
		<td>
			<!--[if (gte mso 9)|(IE)]>
<table width="600" align="center" cellpadding="0" cellspacing="0" border="0">
<tr>
  <td>
<![endif]-->
			<table bgcolor="#ffffff" class="content" align="center" cellpadding="0" cellspacing="0" border="0">
				<tr>
					<td bgcolor="#c7d8a7" class="header">
						<table width="70" align="left" border="0" cellpadding="0" cellspacing="0">
							<tr>
								<td height="70" style="padding: 0 20px 20px 0;">
									<img src="data:image/jpg;base64,R0lGODlhRgBGAPcAAPpWWv///+Px/OpLT8fYp+tLT+xMT+1MUOfz/PT5/vb6/ulKTuvx4OXy/On0/fv9/+73/ehKTslFSPhVWey9vu5OUutMUPRSVvpZXf3+//T6/vf7/ueHd/VjYv/7/PpYXfL4/vlWWvmipPr8+PpYXMlGSO9OUt7ozPlwdPz+/+jv2s7GnMivje1NUf/7+/r8/uzSzP739/vY2fvc3fhbXcfQoulLT/75+clWU8fPofh1eNfjwPNRVc9GSfrT1Pz9//j7/vhydev1/e9laPvf3/dzd+yfpuVYV/pXWvBmaexNUPnP0PNSVu6fpfWfoe5jZu1PU9uAcPBoa+Pt+P76+ur0/epPUvve3+mXnvlvcvvd3u9rbeXs9vifovOTleuFiOtOUtxKTe9rb/L5/s/FnPmkpvTV2euWnedQUfimqPpaXuqgqPOUl8zDmu1RVNBVVOXL1fJXWfva2vvZ2vSXmetlavSlp/75+vXa3fh0d+iaofP5/uXQ2unv+O1OUu5iZepMUMh/bfRvdNBUUu5pbf729vWjpePs98h+bfipq/va2+5tcPrR0vX6/vWZm/vU1dGujehLTv77++xPUvb7/uTx/PhzdthHS+nR2v7+///8/PH4/vj8/vioqvTZ3fhUWNaxkMpFSOnt9uXu+ORKTu6IigAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACH/C1hNUCBEYXRhWE1QPD94cGFja2V0IGJlZ2luPSLvu78iIGlkPSJXNU0wTXBDZWhpSHpyZVN6TlRjemtjOWQiPz4gPHg6eG1wbWV0YSB4bWxuczp4PSJhZG9iZTpuczptZXRhLyIgeDp4bXB0az0iQWRvYmUgWE1QIENvcmUgNS41LWMwMTQgNzkuMTUxNDgxLCAyMDEzLzAzLzEzLTEyOjA5OjE1ICAgICAgICAiPiA8cmRmOlJERiB4bWxuczpyZGY9Imh0dHA6Ly93d3cudzMub3JnLzE5OTkvMDIvMjItcmRmLXN5bnRheC1ucyMiPiA8cmRmOkRlc2NyaXB0aW9uIHJkZjphYm91dD0iIiB4bWxuczp4bXA9Imh0dHA6Ly9ucy5hZG9iZS5jb20veGFwLzEuMC8iIHhtbG5zOnhtcE1NPSJodHRwOi8vbnMuYWRvYmUuY29tL3hhcC8xLjAvbW0vIiB4bWxuczpzdFJlZj0iaHR0cDovL25zLmFkb2JlLmNvbS94YXAvMS4wL3NUeXBlL1Jlc291cmNlUmVmIyIgeG1wOkNyZWF0b3JUb29sPSJBZG9iZSBQaG90b3Nob3AgQ0MgKE1hY2ludG9zaCkiIHhtcE1NOkluc3RhbmNlSUQ9InhtcC5paWQ6MzM4QTMwMDM1MzIwMTFFM0E1RkNBMzY2RjNERDgyQjYiIHhtcE1NOkRvY3VtZW50SUQ9InhtcC5kaWQ6MzM4QTMwMDQ1MzIwMTFFM0E1RkNBMzY2RjNERDgyQjYiPiA8eG1wTU06RGVyaXZlZEZyb20gc3RSZWY6aW5zdGFuY2VJRD0ieG1wLmlpZDozMzhBMzAwMTUzMjAxMUUzQTVGQ0EzNjZGM0REODJCNiIgc3RSZWY6ZG9jdW1lbnRJRD0ieG1wLmRpZDozMzhBMzAwMjUzMjAxMUUzQTVGQ0EzNjZGM0REODJCNiIvPiA8L3JkZjpEZXNjcmlwdGlvbj4gPC9yZGY6UkRGPiA8L3g6eG1wbWV0YT4gPD94cGFja2V0IGVuZD0iciI/PgH//v38+/r5+Pf29fTz8vHw7+7t7Ovq6ejn5uXk4+Lh4N/e3dzb2tnY19bV1NPS0dDPzs3My8rJyMfGxcTDwsHAv769vLu6ubi3trW0s7KxsK+urayrqqmop6alpKOioaCfnp2cm5qZmJeWlZSTkpGQj46NjIuKiYiHhoWEg4KBgH9+fXx7enl4d3Z1dHNycXBvbm1sa2ppaGdmZWRjYmFgX15dXFtaWVhXVlVUU1JRUE9OTUxLSklIR0ZFRENCQUA/Pj08Ozo5ODc2NTQzMjEwLy4tLCsqKSgnJiUkIyIhIB8eHRwbGhkYFxYVFBMSERAPDg0MCwoJCAcGBQQDAgEAACH5BAAAAAAALAAAAABGAEYAAAj/AAkIFHhiRICDCBMqXMiwocOHAUacGEhR4A6IGDNqfLij4kAVG0OKxKjCo0AGI1OqRMjAJAGUK2OGbGkSpsybEGl6tImz50KdFXn6HBoAKEWhRHsaHYj0YIYHUKNKnTo1Q9KlJxm+qCSgq9evYMNuIIr1JcNGYdOqFQCCrMumAcZAmEu3rl27mzK5rZn06tu+e3cCHloW7uCUhQ8r/av4ZmKtCiJLnky5suXLlSlZ/WmyBgyGQNaKHk26qwOGMGpUZIGDAsMNpWPL7oqAIQUcLATmQFRCguvXCYILH068uPHjxTX8sC2hRKAcbyRI/904JgXpEgaFCDO9uszroS5F/wAAwE0P6gkzCEHAvr379/Djy4+/x3YPUgsWkAcQopSmhRrMJuBoDSzkwRc2DDCAfvsBoAMRCqVQxXwUVmghAm0lpMUiCnbYIHlqJOKBdxq5YIcVHXr4IXk6XEEiRDOIUUABKSq44n4YdDIiJ8j16OOPGqTggiFgzEhjjTc2mAceAzaZFiZbGGCAkUeqmCR5GKwxhZNOHqLHJFJOSWWKV35YhCeYpanmZWYQckCYUo5pZZlYpiHJYXc4AcUBfMIp5oxz0kmeJTMAJoMUfCbq55+BCgrAB0aM0sCklFZq6aWYcoGFH4l26qecjt4oCB9cggVHHZ2m+maYoIa64gdN9P/hwKy01mrrrQ6IckYLqqb6KaADuJpkEIqo9EgSvar665HCJolBGVSEFIMjvCbrK5xjNnslCnJo5MMQ1va6LI3aXomBCDc8FAMdShgQrrLYGhlsudvO0RAjT8D5rqfxAkpvmSSgm1AhbFhQwKL7LirvvP9emYUMBy3xB5UIW6vwwg3TSUIXXgAyAJUHL+ouvxcDm7Ggn1SgIMgit9yvyScLysORLLtcMrAMx1zmBCp/XLPNfy6cs85l8mCBzyCHfLPQQxN9Jc8rJy110kg6HSoTRyM9NdVVW+0o1B1uzTSZXrt6QdY1ylkj2WWHCvbacK/ddrNnx223jXMLO4EJd8NrzWDewtbdt4L5Aa7t3oPn97fhgaNdo+KFM354BZBXrrjk9F6woOX5jYd5uTxzHkEESHye+eajjx5JB6bTy3Pqox/BQev/XmBD6lGQQQPtrlcQARptEADK7ryXGwckA63AAevFh9oBBysIFBAAOw==">
								</td>
							</tr>
						</table>
						<!--[if (gte mso 9)|(IE)]>
	<table width="425" align="left" cellpadding="0" cellspacing="0" border="0">
	  <tr>
		<td>
  <![endif]-->
						<table class="col425" align="left" border="0" cellpadding="0" cellspacing="0"
							style="width: 100%; max-width: 425px;">
							<tr>
								<td height="70">
									<table width="100%" border="0" cellspacing="0" cellpadding="0">
										<tr>
											<td class="subhead" style="padding: 0 0 0 5px;">
												时间：{{.T}}
											</td>
										</tr>
										<tr>
											<td class="h1" style="padding: 5px 0 0 0;">
												{{.Title}}
											</td>
										</tr>
									</table>
								</td>
							</tr>
						</table>
						<!--[if (gte mso 9)|(IE)]>
		</td>
	  </tr>
  </table>
  <![endif]-->
					</td>
				</tr>
				<tr>
					<td class="innerpadding borderbottom">
						<table width="100%" border="0" cellspacing="0" cellpadding="0">
							<tr>
								<h2> {{ .Message }} </h2>
							</tr>
						</table>
					</td>
				</tr>
				<tr>
					<td class="innerpadding borderbottom">
						<table width="100%" border="0" cellspacing="0" cellpadding="0">
							{{ if .Stack }}
							<tr>
								堆栈信息：
							</tr>
							<tr>
								<td class="bodycopy">
									{{ range $key, $value := .Stack }}
									<strong>{{ $key }}</strong>:<br>
									&nbsp;&nbsp;&nbsp;&nbsp;{{ $value }}<br>
									{{ end }}
								</td>
							</tr>
							{{end}}
						</table>
					</td>
				</tr>
				<tr>
					{{if .Detail}}
					<td class="innerpadding borderbottom">
						<table width="115" align="left" border="0" cellpadding="0" cellspacing="0">
							<tr>
								<td height="115" style="padding: 0 20px 20px 0;">
									详细信息
								</td>
							</tr>
						</table>
						<!--[if (gte mso 9)|(IE)]>
	<table width="380" align="left" cellpadding="0" cellspacing="0" border="0">
	  <tr>
		<td>
  <![endif]-->
						<table class="col380" align="left" border="0" cellpadding="0" cellspacing="0"
							style="width: 100%; max-width: 380px;">
							<tr>
								<td>
									<table width="100%" border="0" cellspacing="0" cellpadding="0">
										{{ range $key, $value := .Detail }}
										<tr>
											<li><strong>{{ $key }}</strong>: {{ $value }}</li>
										</tr>
										{{ end }}
									</table>
								</td>
							</tr>
						</table>
						<!--[if (gte mso 9)|(IE)]>
		</td>
	  </tr>
  </table>
  <![endif]-->
					</td>
					{{end}}
				</tr>
				<tr>
					<!-- <td class="innerpadding borderbottom">
				
				</td> -->
				</tr>
				<tr>
					<!-- <td class="innerpadding bodycopy">

				</td> -->
				</tr>
				<tr>
					<td class="footer" bgcolor="#44525f">
						<table width="100%" border="0" cellspacing="0" cellpadding="0">
							<tr>
								<td align="center" class="footercopy">
									<!-- &reg; Someone, somewhere 20XX<br />
								<a href="#" class="unsubscribe">
									<font color="#ffffff">Unsubscribe</font>
								</a> -->
									<span class="hide">请注意，该邮件地址不接收回复邮件。</span>
								</td>
							</tr>
						</table>
					</td>
				</tr>
			</table>
			<!--[if (gte mso 9)|(IE)]>
  </td>
</tr>
</table>
<![endif]-->
		</td>
	</tr>
</table>
</body>

</html>`
