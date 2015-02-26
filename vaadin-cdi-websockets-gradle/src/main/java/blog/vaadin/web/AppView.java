package blog.vaadin.web;

import com.vaadin.cdi.CDIView;
import com.vaadin.navigator.View;
import com.vaadin.navigator.ViewChangeListener.ViewChangeEvent;
import com.vaadin.ui.Component;
import com.vaadin.ui.Label;
import com.vaadin.ui.Panel;
import com.vaadin.ui.VerticalLayout;

@CDIView
public class AppView extends Panel implements View {

    @Override
    public void enter(ViewChangeEvent event) {
        setSizeFull();
        Component content = buildContent();
        setContent(content);
    }

    protected Component buildContent() {
        VerticalLayout layout = new VerticalLayout();
        layout.setSizeFull();
        Label h1 = new Label("Vaadin App");
        h1.addStyleName("h1");
        layout.addComponent(h1);
        return layout;
    }

}
