var fs = require('fs');
var console = require('better-console');
var markdown = require('markdown').markdown;
var args = process.argv;

if(args.length !== 4) {
    console.log('need 2 arguments');
    process.exit(0);
}

var mdfile = args[2];
var tarfile = args[3];

fs.readFile(mdfile, function(err, data) {
    if(err) {
        console.error('read ' + mdfile, tarfile, err);
        process.exit(1);
    }
    var htext = markdown.toHTML(data.toString());
    fs.readFile('./template.html', function(err, tem) {
        if(err) {
            console.error('read template.html', tarfile, err);
            process.exit(1);
        }
        var temtext = tem.toString(); 
        var final = temtext.replace(/CONTENT/g, htext);
        fs.writeFile(tarfile, final, function(err){
            if(err) {
                console.error('write error', tarfile, err);
                process.exit(1);
            }
            console.info('write success. source:', mdfile, ', target:', tarfile);
        });
    });
});
