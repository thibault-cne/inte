# Import needed packages


# Import personal packages
from App import create_app


# Start app if file is not imported
if __name__ == "__main__":
    app = create_app()
    app.run(debug=True, host='0.0.0.0', port=5454)
