describe('The Home Page', () => {
  it('successfully loads', () => {
    cy.visit('http://localhost:3222');
    cy.screenshot();
  })
})
