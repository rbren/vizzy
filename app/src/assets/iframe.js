var allowedOrigins = ['http://localhost:3030', 'https://vizzy.rbren.io'];

window.d3Intervals = [];
function clearAllTimeouts() {
  var id = window.setTimeout(function() {}, 0);
  while (id--) {
    window.clearTimeout(id);
  }
  id = window.setInterval(function() {}, 0);
  while (id--) {
    window.clearInterval(id);
  }
  id = window.requestAnimationFrame(function() {});
  while (id--) {
    window.cancelAnimationFrame(id);
  }
  window.d3Intervals.forEach(function(interval) {
    interval.stop();
  });
  window.d3Intervals = [];
}

// hack d3 interval to make sure they get cleared
var d3Interval = d3.interval;
d3.interval = function(callback, delay, time) {
  var interval = d3Interval(callback, delay, time);
  window.d3Intervals.push(interval);
  return interval;
}

function extractErrorMessage(e) {
  if (typeof e === 'string') {
    return e;
  } else if (e instanceof Error) {
    var source = (e.stack.split('\n').find(l => l.includes('anonymous')) || '').replaceAll(/^.*(<anonymous>:\d+:\d+).*$/g, '$1');
    if (source) {
      return e.message + ' at ' + source;
    } else {
      return e.message;
    }
  } else {
    return 'Unknown error';
  }
}

window.rerun = _.debounce(function() {
  clearAllTimeouts();
  try {
    eval(window.vizzy.code);
  } catch (e) {
    returnMessage({error: e});
  }
  if (typeof drawVisualization !== 'undefined') {
    window.drawVisualization = drawVisualization;

    var svg = d3.select('svg');
    svg.selectAll('*').remove();
    svg.attr('height', window.innerHeight - window.vizzy.heightOffset);
    svg.attr('width', window.innerWidth - window.vizzy.widthOffset);
    try {
      window.drawVisualization(svg, window.vizzy.data).then(function() {
        returnMessage({success: true});
      }).catch(function(e) {
        returnMessage({error: e});
      });
    } catch (e) {
      returnMessage({error: e});
      throw e;
    }
    checkSVGForIssues();
  } else if (typeof computeMetadata !== 'undefined') {
    try {
      var metadata = computeMetadata(window.vizzy.data);
    } catch (e) {
      returnMessage({error: e});
      return;
    }
    returnMessage({metadata: metadata});
  }
}, 500);

window.returnMessage = function(msg) {
  msg.id = window.vizzy.id;
  if (msg.error) msg.error = extractErrorMessage(msg.error);
  try {
    window.parent.postMessage(msg, window.vizzy.origin || '');
  } catch (e) {
    console.log('Failed to post message to parent window', e);
    return returnMessage({
      error: 'Failed to execute JavaScript. Please try again.'
    });
  }
}

window.checkSVGForIssues = function() {
  var queries = [
    'path[d*="NaN"]',
    'rect[x="NaN"]',
    'rect[y="NaN"]',
    'rect[width="NaN"]',
    'rect[height="NaN"]',
    'rect[width^="-"]',
    'rect[height^="-"]',
    'circle[cx="NaN"]',
    'circle[cy="NaN"]',
    'circle[r="NaN"]',
    'line[x1="NaN"]',
    'line[y1="NaN"]',
    'line[x2="NaN"]',
    'line[y2="NaN"]',
    'text[x="NaN"]',
    'text[y="NaN"]',
  ];
  queries.forEach(function(query) {
    var elements = document.querySelectorAll(query);
    if (elements.length) {
      var elem = query.split('[')[0];
      var attr = query.split('[')[1].split('=')[0];
      if (attr.endsWith('^') || attr.endsWith('$') || attr.endsWith('*')) {
        attr = attr.slice(0, -1);
      }
      returnMessage({
        error: 'SVG contains a <' + elem + '> with invalid data in the ' + attr + ' attribute',
      });
    }
  });
}

window.setUpEnvironment = function(env) {
  window.vizzy = env;
}

window.addEventListener('message', function(event) {
  if (!allowedOrigins.includes(event.origin)) {
    throw new Error('Invalid origin:' + event.origin);
    return;
  }
  setUpEnvironment({
    data: event.data.data,
    code: event.data.code,
    id: event.data.id,
    origin: event.origin,
    heightOffset: event.data.heightOffset || 0,
    widthOffset: event.data.widthOffset || 0,
  });
  window.rerun();
}, false);

window.addEventListener('resize', function(event) {
  if (window.rerun) {
    window.rerun.cancel();
    window.rerun();
  }
});
