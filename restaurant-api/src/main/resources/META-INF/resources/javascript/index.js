'use strict';

const e = React.createElement;

class LikeButton extends React.Component {
  constructor(props) {
    super(props);
    this.state = { liked: false };
  }

  render() {
    if (this.state.liked) {
      return 'You liked this.';
    }

    return e(
      'button',
      { onClick: () => this.setState({ liked: true }) },
      'Like'
    );
  }
}

class Header extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
    return (
      <div className="banner lead">
        {this.props.name}
      </div>
    );
  }
}

class Courses extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
    return (
      <React.Fragment>
        <h2>{this.props.name}</h2>
        {
          this.props.courses.map((course, i) => {
            return (
              <p key={i}>
                <span className="course-name">{course.name}</span>:
                <span className="course-desc">{course.description}</span> -
                <span className="course-price"> {course.price} €</span>
              </p>
            );
          })
        }
      </React.Fragment>
    );
  }
}

class InfoBadge extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
    return (
      <div className="right-column">
        <div className="right-section">
          <h3>About:</h3>
          <ul>
            <li><span className="bold">Name:</span> {this.props.name}</li>
            <li><span className="bold">Food:</span> {this.props.foodType}</li>
            {
                this.props.address ? <li><span className="bold">Address:</span> {this.props.address}</li> : ""
            }
            {
                this.props.phoneNumber ? <li><span className="bold">Reservations:</span> {this.props.phoneNumber}</li> : ''
            }
          </ul>
        </div>
      </div>
    );
  }
}

class Restaurant extends React.Component {
  constructor(props) {
    super(props);
    this.state = {};
    fetch("/api")
      .then(r => r.json())
      .then(r => {
        console.log(r);
        this.setState({
          name: r.name,
          foodType: r.foodType,
          address: r.address,
          phoneNumber: r.phoneNumber
        });
      });
    fetch("/api/menu")
      .then(r => r.json())
      .then(menu => {
        console.log(menu);
        this.setState({ menu: menu });
      })
  }

  render() {
    return (
      <React.Fragment>
        <Header name={this.state.name} />
        <div className="container">
          <div className="left-column">
            <p className="lead">The menu</p>
            {this.state.menu && <Courses name="Starters" courses={this.state.menu.starters}></Courses>}
            {this.state.menu && <Courses name="Main course" courses={this.state.menu.main}></Courses>}
            {this.state.menu && <Courses name="Desserts" courses={this.state.menu.desserts}></Courses>}
          </div>
          <InfoBadge name={this.state.name} address={this.state.address} foodType={this.state.foodType} phoneNumber={this.state.phoneNumber}></InfoBadge>
        </div>
      </React.Fragment>
    );
  }
}

const domContainer = document.getElementById("main");
ReactDOM.render(e(Restaurant), domContainer);