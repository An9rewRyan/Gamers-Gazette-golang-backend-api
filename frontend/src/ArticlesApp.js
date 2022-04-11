import './App.css';
import React from 'react';

const PORT = process.env.PORT || 3000
console.log(PORT)
class ArticlesApp extends React.Component {
   
  constructor(props) {
      super(props);
 
      this.state = {
          items: [],
          DataisLoaded: false
      };
  }

  componentDidMount() {
      fetch(
`https://gamersgazette.herokuapp.com/api/v1/articles/`)
          .then((res) => res.json())
          .then((json) => {
              this.setState({
                  items: json,
                  DataisLoaded: true
              });
          })
  }
  render() {
      const { DataisLoaded, items } = this.state;
      if (!DataisLoaded) return <div>
          <h1> Pleses wait some time.... </h1> </div> ;
 
      return (
      <div className = "App">
          <h1> Fetch data from an api in react </h1>  {
              items.map((item) => ( 
              <ol key = { item.article_id } >
                  Title: { item.title }, 
                  Pub_date: { item.pub_date },
                  Image_url: {item.image_url},
                  Content: { item.content }, 
                  Src_link: {item.src_link},
                  Site_alias: {item.site_alias}
                  </ol>
              ))
          }
      </div>
  );
}
}

export default ArticlesApp;