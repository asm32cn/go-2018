#!/usr/bin/env python
# -*- coding: utf-8 -*-

# Python 2.7
# asm32-buff-sqlite3-web-py.py

import web, datetime, os

_folder = 'E:\\gitlearn\\2\\python-2018\\py-2018-dev\\asm32-buff-sqlite3-web-py.py\\f\\'

_template = '''
<!doctype html>
<html lang="en">
<head>
	<meta charset="UTF-8">
    <title>PA_div_contenteditable.html</title>
    <meta http-equiv="pragma" content="no-cache" />
    <meta HTTP-EQUIV="Cache-Control" CONTENT="no-cache, must-revalidate"/>
    <meta HTTP-EQUIV="expires" CONTENT="-1"/>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
    <meta http-equiv="X-UA-Compatible" content="IE=100" >
</head>
<body>

<p>
	<button>精简</button>
	<button>保存</button>
</p>

<form method="post">
	<p>
		<input name="txtTitle" size="150"/>
	<p>
	<p>
		<textarea name="txtContent" rows="5" cols="180"></textarea>
	<p>
</form>

<div contenteditable="true" style="width:100%; height:1600px;border:1px solid #CCC;"></div>

<script type="text/javascript">
<!--
window.onload = function(){
	var buttons = document.getElementsByTagName('button');
	var m_objDisplay = document.getElementsByTagName('textarea')[0];
	var m_objContent = document.getElementsByTagName('div')[0];
	var _form = document.getElementsByTagName('form')[0];

	buttons[0].onclick = function(){
		var m_strContent = m_objContent.innerHTML;
		m_strContent = m_strContent.replace(/ style=".*?"/gi, '');
		m_strContent = m_strContent.replace(/ class=".*?"/gi, '');
		m_strContent = m_strContent.replace(/ height=".*?"/gi, '');
		m_strContent = m_strContent.replace(/<[\/]?span.*?>/gi, '');
		m_strContent = m_strContent.replace(/<p .*?>/gi, '<p>');
		m_objContent.innerHTML = m_strContent;
		alert('Done.');
	}

	buttons[1].onclick = function(){
		var m_strContent = m_objContent.innerHTML;
		m_objDisplay.value = m_objContent.innerHTML;
		// m_objContent.innerHTML = m_strContent;
		if(confirm('Save ?')){
			// alert('Save.');
			_form.submit()
		}
	}
}
//-->
</script>
</body>
</html>
'''

urls = (
	'/', 'index'
)

if not os.path.exists(_folder):os.makedirs(_folder)

def _save(strTitle, strContent):
	dt = datetime.datetime.now() # datetime.datetime(2016, 5, 22, 3, 17, 3, 203000)
	_filename = _folder + '%04d-%02d-%02d-%02d-%02d-%02d.txt' % (dt.year, dt.month, dt.day, dt.hour, dt.minute, dt.second)
	print 'Save..'
	with open(_filename, 'w') as fo:
		fo.write( (strTitle + '\n\n' + strContent).encode('utf-8') )

class index():
	def GET(self):
		# return 'asm32-buff-sqlite3-web-py.py.GET'
		return _template

	def POST(self):
		_input = web.input()
		strTitle = _input.get('txtTitle')
		strContent = _input.get('txtContent')

		_save(strTitle, strContent)

		print 'strTitle=' + strTitle
		return 'asm32-buff-sqlite3-web-py.py.POST'

if __name__ == '__main__':
	web.application(urls, globals()).run()
