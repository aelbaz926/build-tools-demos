from flask import Flask, jsonify
import requests

app = Flask(__name__)

@app.route('/')
def home():
    return jsonify({
        'message': 'pip Demo Application',
        'tool': 'pip',
        'language': 'Python',
        'framework': 'Flask',
        'version': '1.0.0'
    })

@app.route('/info')
def info():
    return jsonify({
        'dependencies': {
            'Flask': '3.0.0',
            'requests': '2.31.0'
        },
        'description': 'REST API demonstrating pip package management',
        'endpoints': ['/', '/info', '/external']
    })

@app.route('/external')
def external():
    try:
        response = requests.get('https://pypi.org/pypi/pip/json')
        data = response.json()
        return jsonify({
            'message': 'External API call successful',
            'pipInfo': {
                'name': data['info']['name'],
                'version': data['info']['version'],
                'summary': data['info']['summary']
            }
        })
    except Exception as e:
        return jsonify({'error': str(e)}), 500

if __name__ == '__main__':
    print('=== pip Demo Application ===\n')
    print('✅ Server running on http://localhost:5000')
    print('\nAvailable endpoints:')
    print('  GET http://localhost:5000/')
    print('  GET http://localhost:5000/info')
    print('  GET http://localhost:5000/external')
    print('\nPress Ctrl+C to stop\n')
    
    app.run(debug=True, port=5000)
