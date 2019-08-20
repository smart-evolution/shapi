describe('Login page', () => {
  it('should render successfully', () => {
    cy.visit(Cypress.env('SH_CYPRESS_URL'));
    cy.screenshot();
  });

  it('should navigate to registration', () => {
    cy.get('.tst-signup').click();
    cy.screenshot();
  });

  it('should register successfully', () => {
    cy.get('.tst-email').type('admin@localhost.pl');
    cy.get('.tst-password').type('admin');
    cy.get('.tst-register').click();
  });

  it('should navigate back to login page', () => {
    cy.get('.tst-login').click();
    cy.screenshot();
  });

  it('should login successfully', () => {
    cy.get('.tst-email').type('admin@localhost.pl');
    cy.get('.tst-password').type('admin');
    cy.get('.tst-login').click();
    cy.screenshot();
  });
});
