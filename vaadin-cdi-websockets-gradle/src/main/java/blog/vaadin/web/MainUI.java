package blog.vaadin.web;

import javax.inject.Inject;

import com.vaadin.annotations.Push;
import com.vaadin.annotations.Theme;
import com.vaadin.cdi.CDIUI;
import com.vaadin.cdi.CDIViewProvider;
import com.vaadin.navigator.Navigator;
import com.vaadin.server.VaadinRequest;
import com.vaadin.ui.UI;
import com.vaadin.ui.themes.ValoTheme;

@CDIUI("")
@Push
@Theme(ValoTheme.THEME_NAME)
public class MainUI extends UI {

    @Inject
    CDIViewProvider viewProvider;

    @Override
    protected void init(VaadinRequest request) {
        setSizeFull();
        Navigator navigator = new Navigator(this, this);
        navigator.addProvider(viewProvider);
        if ("".equals(navigator.getState())) {
            navigator.navigateTo("app");
        }
    }
}
