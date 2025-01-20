import React from 'react';

const Header: React.FC = () => {
  return (
    <header style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', padding: '1rem 2rem', borderBottom: '1px solid #ddd' }}>
      <div>
        <h1 style={{ margin: 0 }}>My Vocabs</h1>
      </div>
      <nav style={{ display: 'flex', gap: '1rem' }}>
        {/* <a href="/contact">Contact</a> */}
        <button style={{ border: '1px solid #000', background: 'transparent', padding: '0.5rem 1rem', cursor: 'pointer' }}>Sign In</button>
        <button style={{ background: '#000', color: '#fff', padding: '0.5rem 1rem', border: 'none', cursor: 'pointer' }}>Register</button>
      </nav>
    </header>
  );
};

export default Header;
