import QtQuick 2.13
import QtQuick.Controls 2.13

import "../."
import "../controls"

Pane {
    id: root

    padding: 0
    topPadding: 10

    Column {
        spacing: 10

        CheckBox {
            text: qsTr("Use plain-text HTTP/2 when connecting to server (no TLS)")
            checked: wc.options.plaintext
            onCheckedChanged: wc.options.plaintext = checked
        }

        CheckBox {
            text: qsTr("Skip server certificate and domain verification (insecure)")
            checked: wc.options.insecure
            onCheckedChanged: wc.options.insecure = checked 
        }

        TextAreaField {
            labelText: qsTr("Trusted root certificate(s)")
            text: wc.options.rootca
            onTextChanged: wc.options.rootca = text 
        }

        Row {
            spacing: 10
            TextAreaField {
                labelText: qsTr("Client certificate (public key)")
                text: wc.options.clientcert
                onTextChanged: wc.options.clientcert = text
            }

            TextAreaField {
                labelText: qsTr("Client private key")
                text: wc.options.clientkey
                onTextChanged: wc.options.clientkey = text
            }
        }
    }
}


