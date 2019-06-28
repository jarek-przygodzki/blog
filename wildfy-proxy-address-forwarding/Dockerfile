FROM jboss/wildfly:16.0.0.Final

RUN set -x && \
    "$JBOSS_HOME/bin/jboss-cli.sh" --commands="embed-server,  /subsystem=undertow/server=default-server/http-listener=default:write-attribute(name=proxy-address-forwarding,value=true)" && \
    # Fix for Error: Could not rename /opt/jboss/wildfly/standalone/configuration/standalone_xml_history/current
    rm -rf /opt/jboss/wildfly/standalone/configuration/standalone_xml_history

ADD build/libs/proxy-address-forwarding-demo.war /opt/jboss/wildfly/standalone/deployments/
