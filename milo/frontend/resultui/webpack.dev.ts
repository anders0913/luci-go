// Copyright 2020 The LUCI Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import fs from 'fs';
import path from 'path';

import { createProxyMiddleware } from 'http-proxy-middleware';
import webpack from 'webpack';
import merge from 'webpack-merge';

import common from './webpack.common';

const config: webpack.Configuration = merge(common, {
  mode: 'development',
  devtool: 'eval-source-map',

  devServer: {
    contentBase: path.join(__dirname, './out/'),
    historyApiFallback: true,
    https: {
      key: fs.readFileSync(path.join(__dirname, 'dev-configs/cert.key')),
      cert: fs.readFileSync(path.join(__dirname, 'dev-configs/cert.pem')),
    },
    before: (app) => {
      app.get('/configs.js', async (_req, res) => {
        res.set('content-type', 'application/javascript');
        const appConfigs = JSON.parse(fs.readFileSync('./dev-configs/configs.json', 'utf8'));
        const configsTemplate = fs.readFileSync('./configs.template.js', 'utf8');
        const config = configsTemplate
          .replace('{{.ResultDB.Host}}', appConfigs.RESULT_DB.HOST)
          .replace('{{.Buildbucket.Host}}', appConfigs.BUILDBUCKET.HOST)
          .replace('{{.OAuth2.ClientID}}', appConfigs.OAUTH2.CLIENT_ID);
        res.send(config);
      });

      app.use(/^(?!\/(ui|static\/(dist|styles))\/).*/, createProxyMiddleware({
        // This attribute is required. However the value will be overridden by
        // the router option. So the value doesn't matter.
        target: 'https://luci-milo-dev.appspot.com',
        router: () => {
          const localDevConfigs = JSON.parse(fs.readFileSync('./dev-configs/local-dev-configs.json', 'utf8'));
          return localDevConfigs.milo.url;
        },
        changeOrigin: true,
      }));
    },
  },
});

// Default export is required by webpack.
// tslint:disable-next-line: no-default-export
export default config;
